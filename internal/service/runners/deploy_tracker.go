package runners

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/book-svc/internal/eth_reader"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

const deployTrackerCursor = "deploy_tracker_last_block"

var (
	BookNotFoundErr     = errors.New("book with specified filters was not found")
	TxStatusNotFoundErr = errors.New("tx status was not found")
)

type DeployTracker struct {
	log      *logan.Entry
	database data.DB
	rpc      *ethclient.Client
	reader   eth_reader.FactoryContractReader

	name          string
	address       common.Address
	firstBlock    uint64
	iterationSize uint64
	runnerCfg     config.Runner
}

func NewDeployTracker(cfg config.Config) *DeployTracker {
	return &DeployTracker{
		log:      cfg.Log(),
		database: postgres.NewDB(cfg.DB()),
		rpc:      cfg.EtherClient().Rpc,
		reader:   eth_reader.NewFactoryContractReader(cfg.EtherClient().Rpc),

		name:          cfg.DeployTracker().Name,
		address:       cfg.DeployTracker().Address,
		firstBlock:    cfg.DeployTracker().FirstBlock,
		iterationSize: cfg.DeployTracker().IterationSize,
		runnerCfg:     cfg.DeployTracker().Runner,
	}
}

func (t *DeployTracker) Run(ctx context.Context) {
	running.WithBackOff(
		ctx,
		t.log,
		t.name,
		t.Track,
		t.runnerCfg.NormalPeriod,
		t.runnerCfg.MinAbnormalPeriod,
		t.runnerCfg.MaxAbnormalPeriod,
	)
}

func (t *DeployTracker) Track(ctx context.Context) error {
	previousBlock, err := t.GetPreviousBlock()
	if err != nil {
		return errors.Wrap(err, "failed to get previous block")
	}

	must, err := t.MustNotExceedLastBlock(previousBlock)
	if err != nil {
		return errors.Wrap(err, "failed to check whether previous block is less than the last block in chain")
	}
	if !must {
		return nil
	}

	t.log.Debugf("Trying to iterate from block %d to %d...", previousBlock, previousBlock+t.iterationSize)

	events, _, err := t.reader.GetDeployEvents(t.address, previousBlock, previousBlock+t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get events")
	}

	for _, event := range events {
		t.log.Debugf("Caught new deploy event with block number %d", event.BlockNumber)

		if err = t.ProcessEvent(event); err != nil {
			return errors.Wrap(err, "failed to insert event into the database")
		}

		t.log.Debugf("Successfully inserted contract into the database")
	}

	newBlock, err := t.GetNewBlock(previousBlock, t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get new block")
	}

	t.log.Debugf("New block value is %d", newBlock)

	if err = t.database.KeyValue().Upsert(data.KeyValue{
		Key:   deployTrackerCursor,
		Value: strconv.FormatInt(newBlock, 10),
	}); err != nil {
		return errors.Wrap(err, "failed to upsert new value")
	}

	t.log.Debugf("Updated KV cursor value")
	return nil
}

func (t *DeployTracker) MustNotExceedLastBlock(block uint64) (bool, error) {
	// Retrieving the last blockchain block number
	lastBlockchainBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "failed to get the last block in the blockchain")
	}

	return block <= lastBlockchainBlock, nil
}

func (t *DeployTracker) GetPreviousBlock() (uint64, error) {
	cursorKV, err := t.database.KeyValue().Get(deployTrackerCursor)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get cursor value")
	}
	if cursorKV == nil {
		t.log.Debug("Empty key value cursor, setting 0")
		cursorKV = &data.KeyValue{
			Key:   deployTrackerCursor,
			Value: "0",
		}
	}

	cursor, err := strconv.ParseInt(cursorKV.Value, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert cursor value from string to integer")
	}
	t.log.Debugf("Cursor value is %d", cursor)

	cursorUInt64 := uint64(cursor)
	if cursorUInt64 > t.firstBlock {
		t.log.Debugf("Cursor has a greater value than a config one. Choosing cursor value")
		return cursorUInt64, nil
	}

	t.log.Debugf("Config value has a greater value than a cursor one. Choosing config value")
	return t.firstBlock, nil
}

func (t *DeployTracker) GetNewBlock(previousBlock, iterationSize uint64) (int64, error) {
	// Retrieving the last blockchain block number
	lastBlockchainBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return 0, errors.Wrap(err, "failed to get the last block in the blockchain")
	}

	t.log.Debugf("Last blockchain block has id of %d", lastBlockchainBlock)

	if previousBlock+iterationSize+1 > lastBlockchainBlock {
		return int64(lastBlockchainBlock + 1), nil
	}

	return int64(previousBlock + iterationSize + 1), nil
}

func (t *DeployTracker) ProcessEvent(event eth_reader.DeployEvent) error {
	book, err := t.database.Books().New().FilterByTokenId(int64(event.TokenId)).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get book by token id")
	}
	if book == nil {
		t.log.Warnf("Book with token id %v was not found", event.TokenId)

		return nil
		//return errors.From(BookNotFoundErr, logan.F{
		//	"token_id": event.TokenId,
		//})
	}

	switch event.Status {
	case types.ReceiptStatusSuccessful:
		if err = t.database.Books().UpdateContractAddress(event.Address.String(), book.ID); err != nil {
			return errors.Wrap(err, "failed to update contract address", logan.F{
				"contract_address": event.Address.String(),
			})
		}

		return t.database.Books().UpdateDeployStatus(resources.DeploySuccessful, book.ID)
	case types.ReceiptStatusFailed:
		return t.database.Books().UpdateDeployStatus(resources.DeployFailed, book.ID)
	}

	return errors.From(TxStatusNotFoundErr, logan.F{
		"block_number": event.BlockNumber,
	})
}
