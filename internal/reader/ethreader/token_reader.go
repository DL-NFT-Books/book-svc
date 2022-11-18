package ethreader

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/book-svc/internal/reader"
	"gitlab.com/tokend/nft-books/book-svc/solidity/generated/itokencontract"
)

var NullIteratorErr = errors.New("iterator has a nil value")

type TokenContractReader struct {
	rpc *ethclient.Client

	from    *uint64
	to      *uint64
	address *common.Address
}

func NewTokenContractReader(cfg config.Config) reader.TokenReader {
	return &TokenContractReader{
		rpc: cfg.EtherClient().Rpc,
	}
}

func (r *TokenContractReader) From(from uint64) reader.TokenReader {
	r.from = &from
	return r
}

func (r *TokenContractReader) To(to uint64) reader.TokenReader {
	r.to = &to
	return r
}

func (r *TokenContractReader) WithAddress(address common.Address) reader.TokenReader {
	r.address = &address
	return r
}

func (r *TokenContractReader) validateParameters() error {
	if r.from == nil {
		return reader.FromNotSpecifiedErr
	}
	if r.address == nil {
		return reader.AddressNotSpecifiedErr
	}

	return nil
}

func (r *TokenContractReader) GetUpdateEvents() ([]ethereum.UpdateEvent, error) {
	if err := r.validateParameters(); err != nil {
		return nil, err
	}

	events := make([]ethereum.UpdateEvent, 0)

	instance, err := itokencontract.NewItokencontract(*r.address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize a contract instance")
	}

	iterator, err := instance.FilterTokenContractParamsUpdated(
		&bind.FilterOpts{
			Start: *r.from,
			End:   r.to,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return nil, NullIteratorErr
	}

	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event

		if event != nil {
			events = append(events, ethereum.UpdateEvent{
				Name:        event.TokenName,
				Symbol:      event.TokenSymbol,
				Price:       event.NewPrice.String(),
				BlockNumber: event.Raw.BlockNumber,
			})
		}
	}

	return events, nil
}
