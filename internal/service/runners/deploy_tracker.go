package runners

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/book-svc/internal/reader"
	"gitlab.com/tokend/nft-books/book-svc/internal/reader/ethreader"
	"gitlab.com/tokend/nft-books/book-svc/resources"
	network_connector "gitlab.com/tokend/nft-books/network-svc/connector/api"
	network_resources "gitlab.com/tokend/nft-books/network-svc/resources"
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
	reader   reader.FactoryReader
	cfg      config.DeployTracker

	networker *network_connector.Connector
}

func NewDeployTracker(cfg config.Config) *DeployTracker {
	return &DeployTracker{
		log:      cfg.Log(),
		database: postgres.NewDB(cfg.DB()),
		rpc:      cfg.EtherClient().Rpc,
		reader:   ethreader.NewFactoryContractReader(cfg).WithAddress(cfg.DeployTracker().Address),
		cfg:      cfg.DeployTracker(),

		networker: cfg.NetworkConnector(),
	}
}

func (t *DeployTracker) Run(ctx context.Context) {
	running.WithBackOff(
		ctx,
		t.log,
		t.cfg.Name,
		t.Track,
		t.cfg.Runner.NormalPeriod,
		t.cfg.Runner.MinAbnormalPeriod,
		t.cfg.Runner.MaxAbnormalPeriod,
	)
}

func (t *DeployTracker) Track(ctx context.Context) error {
	// getting networks
	networksResponse, err := t.networker.GetNetworks()
	if err != nil {
		return errors.Wrap(err, "failed to form a list of available networks")
	}

	networks := networksResponse.Data
	for _, network := range networks {
		if err = t.ProcessNetwork(network.Attributes); err != nil {
			return errors.Wrap(err, "failed to process specified network", logan.F{
				"network_name": network.Attributes.Name,
				"chain_id":     network.Attributes.ChainId,
			})
		}
	}

	return nil
}

func (t *DeployTracker) ProcessNetwork(network network_resources.NetworkDetailedAttributes) error {
	// TODO: IMPLEMENT SPECIFIC NETWORK TRACKING

	startBlock, err := t.GetStartBlock()
	if err != nil {
		return errors.Wrap(err, "failed to get start block")
	}

	lastBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get the last block in the blockchain")
	}

	if startBlock > lastBlock {
		t.log.Debugf("Start block is greater than the last blockchain block, omitting")
	}

	t.log.Debugf("Trying to iterate from block %d to %d...", startBlock, startBlock+t.cfg.IterationSize)

	events, err := t.reader.
		From(startBlock).
		To(startBlock + t.cfg.IterationSize).
		GetDeployEvents()
	if err != nil {
		return errors.Wrap(err, "failed to get events")
	}

	if len(events) == 0 {
		t.log.Debug("No deploy events found")
	}

	for _, event := range events {
		t.log.Infof("Caught new deploy event with block number %d", event.BlockNumber)

		if err = t.ProcessEvent(event); err != nil {
			return errors.Wrap(err, "failed to insert event into the database")
		}

		t.log.Info("Successfully inserted contract into the database")
	}

	nextBlock := t.GetNextBlock(startBlock, t.cfg.IterationSize, lastBlock)

	if err = t.database.KeyValue().Upsert(data.KeyValue{
		Key:   deployTrackerCursor,
		Value: strconv.FormatInt(nextBlock, 10),
	}); err != nil {
		return errors.Wrap(err, "failed to upsert new value")
	}

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

func (t *DeployTracker) GetStartBlock() (uint64, error) {
	//TODO: CHANGE IMPLEMENTATION FOR VARIOUS NETWORKS

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

	cursorUInt64 := uint64(cursor)
	if cursorUInt64 > t.cfg.FirstBlock {
		return cursorUInt64, nil
	}

	return t.cfg.FirstBlock, nil
}

func (t *DeployTracker) GetNextBlock(startBlock, iterationSize, lastBlock uint64) int64 {
	if startBlock+iterationSize+1 > lastBlock {
		return int64(lastBlock + 1)
	}

	return int64(startBlock + iterationSize + 1)
}

func (t *DeployTracker) ProcessEvent(event ethereum.DeployEvent) error {
	book, err := t.database.Books().New().FilterByTokenId(int64(event.TokenId)).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get book by token id")
	}
	if book == nil {
		t.log.Warnf("Book with token id %v was not found", event.TokenId)
		return nil
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
