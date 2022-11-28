package ethreader

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/book-svc/internal/reader"
	"gitlab.com/tokend/nft-books/book-svc/solidity/generated/tokenfactory"
)

type FactoryContractReader struct {
	rpc *ethclient.Client

	from    *uint64
	to      *uint64
	address *common.Address

	// contractInstancesCache is a map storing already initialized instances of contracts
	contractInstancesCache map[common.Address]*tokenfactory.Tokenfactory

	// rpcInstancesCache is a map storing already initialized instances of RPC connections
	rpcInstancesCache map[string]*ethclient.Client
}

func NewFactoryContractReader() reader.FactoryReader {
	return &FactoryContractReader{
		contractInstancesCache: map[common.Address]*tokenfactory.Tokenfactory{},
		rpcInstancesCache:      map[string]*ethclient.Client{},
	}
}

func (r *FactoryContractReader) GetRPCInstance(rawURL string) (*ethclient.Client, error) {
	// Searching RPC instance in cache, if not found -- create new and store
	cacheInstance, ok := r.rpcInstancesCache[rawURL]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := ethclient.Dial(rawURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert value into eth client", logan.F{
			"raw_url": rawURL,
		})
	}

	r.rpcInstancesCache[rawURL] = newInstance
	return newInstance, nil

}

func (r *FactoryContractReader) getContractInstance(address common.Address) (*tokenfactory.Tokenfactory, error) {
	// Searching contract instance in cache, if not found -- create new and store
	cacheInstance, ok := r.contractInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := tokenfactory.NewTokenfactory(address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance for given address", logan.F{
			"address": address,
		})
	}

	r.contractInstancesCache[address] = newInstance
	return newInstance, nil
}

func (r *FactoryContractReader) From(from uint64) reader.FactoryReader {
	r.from = &from
	return r
}

func (r *FactoryContractReader) To(to uint64) reader.FactoryReader {
	r.to = &to
	return r
}

func (r *FactoryContractReader) WithAddress(address common.Address) reader.FactoryReader {
	r.address = &address
	return r
}

func (r *FactoryContractReader) WithRPC(rpc *ethclient.Client) reader.FactoryReader {
	r.rpc = rpc
	return r
}

func (r *FactoryContractReader) validateParameters() error {
	if r.from == nil {
		return reader.FromNotSpecifiedErr
	}
	if r.address == nil {
		return reader.AddressNotSpecifiedErr
	}
	if r.rpc == nil {
		return reader.RPCNotSpecifiedErr
	}

	return nil
}

func (r *FactoryContractReader) GetDeployEvents() ([]ethereum.DeployEvent, error) {
	if err := r.validateParameters(); err != nil {
		return nil, err
	}

	events := make([]ethereum.DeployEvent, 0)

	instance, err := r.getContractInstance(*r.address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance")
	}

	iterator, err := instance.FilterTokenContractDeployed(
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
			receipt, err := r.rpc.TransactionReceipt(context.Background(), event.Raw.TxHash)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get tx receipt", logan.F{
					"tx_hash": event.Raw.TxHash.String(),
				})
			}

			events = append(events, ethereum.DeployEvent{
				Address:     event.NewTokenContractAddr,
				BlockNumber: event.Raw.BlockNumber,
				Name:        event.TokenName,
				Symbol:      event.TokenSymbol,
				TokenId:     event.TokenContractId.Uint64(),
				Status:      receipt.Status,
			})
		}

	}

	return events, nil
}
