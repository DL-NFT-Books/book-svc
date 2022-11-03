package runners

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/book-svc/internal/eth_reader"
	"strconv"
)

const updateTrackerKVPage = "update_tracker_page"

type UpdateTracker struct {
	log    *logan.Entry
	rpc    *ethclient.Client
	reader eth_reader.TokenContractReader

	db data.DB

	name          string
	capacity      uint64
	iterationSize uint64
	runnerCfg     config.Runner
}

func NewUpdateTracker(cfg config.Config) *UpdateTracker {
	return &UpdateTracker{
		log:    cfg.Log(),
		rpc:    cfg.EtherClient().Rpc,
		reader: eth_reader.NewTokenContractReader(cfg.EtherClient().Rpc),

		db: postgres.NewDB(cfg.DB()),

		name:          cfg.UpdateTracker().Name,
		iterationSize: cfg.UpdateTracker().IterationSize,
		runnerCfg:     cfg.UpdateTracker().Runner,
	}
}

func (t *UpdateTracker) Run(ctx context.Context) {
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

func (t *UpdateTracker) Track(ctx context.Context) error {
	books, err := t.FormList()
	if err != nil {
		return errors.Wrap(err, "failed to form a list of contracts")
	}

	for _, book := range books {
		if err = t.ProcessBook(book); err != nil {
			return errors.Wrap(err, "failed to process specified book", logan.F{
				"book_id": book.ID,
			})
		}
	}

	return nil
}

func (t *UpdateTracker) ProcessBook(book data.Book) error {
	t.log.Debugf("Processing book with id of %d", book.ID)
	lastBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get last block number")
	}

	if book.LastBlock > lastBlock {
		t.log.Debugf("contract last block exceeded last block in the blockchain")
		return nil
	}
	events, _, err := t.reader.GetUpdateEvents(book.Address(), book.LastBlock, book.LastBlock+t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get events")
	}

	if len(events) == 0 {
		t.log.Debug("No events found")
	}

	for _, event := range events {
		t.log.Debugf("Found event with a block number of %d", event.BlockNumber)

		if err = t.ProcessEvent(event, book.ID); err != nil {
			return errors.Wrap(err, "failed to process event", logan.F{
				"event_block_number": event.BlockNumber,
			})
		}
	}

	newBlock, err := t.GetNewBlock(book.LastBlock, t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get new block", logan.F{
			"current_block": book.LastBlock,
		})
	}

	if err = t.db.Books().UpdateLastBlockById(newBlock, book.ID); err != nil {
		return errors.Wrap(err, "failed to update last block")
	}

	return nil
}

func (t *UpdateTracker) GetNewBlock(previousBlock, iterationSize uint64) (uint64, error) {
	// Retrieving the last blockchain block number
	lastBlockchainBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return 0, errors.Wrap(err, "failed to get the last block in the blockchain")
	}

	t.log.Debugf("Last blockchain block has id of %d", lastBlockchainBlock)

	if previousBlock+iterationSize+1 > lastBlockchainBlock {
		return lastBlockchainBlock + 1, nil
	}

	return previousBlock + iterationSize + 1, nil
}

func (t *UpdateTracker) ProcessEvent(event eth_reader.UpdateEvent, id int64) error {
	return t.db.Transaction(func() error {
		if err := t.db.Books().UpdateContractNameByID(event.Name, id); err != nil {
			return errors.Wrap(err, "failed to update status")
		}

		if err := t.db.Books().UpdatePriceByID(strconv.FormatUint(event.Price, 10), id); err != nil {
			return errors.Wrap(err, "failed to update price by id")
		}

		return nil
	})
}

func (t *UpdateTracker) Select(pageNumber uint64) ([]data.Book, error) {
	cutQ := t.db.Books().Page(pgdb.OffsetPageParams{
		Limit:      t.capacity,
		PageNumber: pageNumber})

	return cutQ.Select()
}

func (t *UpdateTracker) FormList() ([]data.Book, error) {
	pageNumberKV, err := t.db.KeyValue().Get(updateTrackerKVPage)
	if pageNumberKV == nil {
		pageNumberKV = &data.KeyValue{
			Key:   updateTrackerKVPage,
			Value: "0",
		}
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get page number from KV table")
	}

	pageNumber, err := strconv.ParseInt(pageNumberKV.Value, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert page number to an integer format")
	}

	contracts, err := t.Select(uint64(pageNumber))
	if err != nil {
		return nil, errors.Wrap(err, "failed to select contracts from the table")
	}

	if len(contracts) == 0 && pageNumber == 0 {
		t.log.Warn("contracts table is empty")
		return nil, nil
	}

	if len(contracts) == 0 {
		if err = t.db.KeyValue().Upsert(data.KeyValue{
			Key:   updateTrackerKVPage,
			Value: "0",
		}); err != nil {
			return nil, errors.Wrap(err, "failed to update last processed contract")
		}

		return t.FormList()
	}

	if err = t.db.KeyValue().Upsert(data.KeyValue{
		Key:   updateTrackerKVPage,
		Value: strconv.FormatInt(pageNumber+1, 10),
	}); err != nil {
		return nil, errors.Wrap(err, "failed to update last processed contract")
	}

	return contracts, nil
}
