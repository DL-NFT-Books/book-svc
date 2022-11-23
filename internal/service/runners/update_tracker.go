package runners

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/book-svc/internal/reader"
	"gitlab.com/tokend/nft-books/book-svc/internal/reader/ethreader"
)

const updateTrackerKVPage = "update_tracker_page"

type UpdateTracker struct {
	log      *logan.Entry
	rpc      *ethclient.Client
	cfg      config.UpdateTracker
	reader   reader.TokenReader
	database data.DB
}

func NewUpdateTracker(cfg config.Config) *UpdateTracker {
	return &UpdateTracker{
		log:      cfg.Log(),
		cfg:      cfg.UpdateTracker(),
		reader:   ethreader.NewTokenContractReader(cfg), //empty reader, set params when process specified network
		database: postgres.NewDB(cfg.DB()),
	}
}

func (t *UpdateTracker) Run(ctx context.Context) {
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

func (t *UpdateTracker) Track(ctx context.Context) error {
	books, err := t.FormList()
	if err != nil {
		return errors.Wrap(err, "failed to form a list of contracts")
	}

	for _, book := range books {
		// setting specific network params before processing book

		// setting new rpc connection according to network params
		rpc, err := t.reader.GetRPCInstance(book.ChainID)
		if err != nil {
			return errors.Wrap(err, "failed to get rpc connection", logan.F{
				"book_id":  book.ID,
				"chain_id": book.ChainID,
			})
		}
		t.rpc = rpc

		// setting new reader according to new rpc and token address
		t.reader = t.reader.
			WithAddress(book.Address()).
			WithRPC(t.rpc)

		if err = t.ProcessBook(book, ctx); err != nil {
			return errors.Wrap(err, "failed to process specified book", logan.F{
				"book_id": book.ID,
			})
		}
	}

	return nil
}

func (t *UpdateTracker) ProcessBook(book data.Book, ctx context.Context) error {
	t.log.Debugf("Processing book with id of %d", book.ID)

	lastBlock, err := t.rpc.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get last block number")
	}

	if book.LastBlock > lastBlock {
		t.log.Debugf("contract last block exceeded last block in the blockchain. Omitting")
		return nil
	}

	events, err := t.reader.
		From(book.LastBlock).
		To(book.LastBlock + t.cfg.IterationSize).
		WithAddress(book.Address()).
		GetUpdateEvents()
	if err != nil {
		return errors.Wrap(err, "failed to get book update events")
	}

	if len(events) == 0 {
		t.log.Debug("No book update events found")
	}

	for _, event := range events {
		t.log.Debugf("Found update book event with a block number %d", event.BlockNumber)

		if err = t.ProcessEvent(event, book.ID); err != nil {
			return errors.Wrap(err, "failed to process event", logan.F{
				"event_block_number": event.BlockNumber,
			})
		}
	}

	nextBlock := t.GetNextBlock(book.LastBlock, t.cfg.IterationSize, lastBlock)

	if err = t.database.Books().UpdateLastBlock(nextBlock, book.ID); err != nil {
		return errors.Wrap(err, "failed to update last block")
	}

	return nil
}

func (t *UpdateTracker) GetNextBlock(startBlock, iterationSize, lastBlock uint64) uint64 {
	t.log.Debugf("Last blockchain block has id of %d", lastBlock)

	if startBlock+iterationSize+1 > lastBlock {
		return lastBlock + 1
	}

	return startBlock + iterationSize + 1
}

func (t *UpdateTracker) ProcessEvent(event ethereum.UpdateEvent, id int64) error {
	if err := t.database.Books().UpdateContractParams(
		event.Name,
		event.Symbol,
		event.Price,
		id,
	); err != nil {
		return errors.Wrap(err, "failed to update contract params")
	}

	return nil
}

func (t *UpdateTracker) Select(pageNumber uint64) ([]data.Book, error) {
	cutQ := t.database.Books().Page(pgdb.OffsetPageParams{
		Limit:      t.cfg.Capacity,
		PageNumber: pageNumber})

	return cutQ.Select()
}

func (t *UpdateTracker) FormList() ([]data.Book, error) {
	pageNumberKV, err := t.database.KeyValue().Get(updateTrackerKVPage)
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
		if err = t.database.KeyValue().Upsert(data.KeyValue{
			Key:   updateTrackerKVPage,
			Value: "0",
		}); err != nil {
			return nil, errors.Wrap(err, "failed to update last processed contract")
		}

		return t.FormList()
	}

	if err = t.database.KeyValue().Upsert(data.KeyValue{
		Key:   updateTrackerKVPage,
		Value: strconv.FormatInt(pageNumber+1, 10),
	}); err != nil {
		return nil, errors.Wrap(err, "failed to update last processed contract")
	}

	return contracts, nil
}
