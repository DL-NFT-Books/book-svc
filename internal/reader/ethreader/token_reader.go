package ethreader

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/book-svc/internal/reader"
	"gitlab.com/tokend/nft-books/book-svc/solidity/generated/itokencontract"
	networkConnector "gitlab.com/tokend/nft-books/network-svc/connector/api"
)

var NullIteratorErr = errors.New("iterator has a nil value")

type TokenContractReader struct {
	rpc *ethclient.Client

	from    *uint64
	to      *uint64
	address *common.Address

	// contractInstancesCache is a map storing already initialized instances of contracts
	contractInstancesCache map[common.Address]*itokencontract.Itokencontract

	// rpcInstancesCache is a map storing already initialized instances of RPC connections
	rpcInstancesCache map[int64]*ethclient.Client

	networker *networkConnector.Connector
}

func NewTokenContractReader(cfg config.Config) reader.TokenReader {
	return &TokenContractReader{
		contractInstancesCache: map[common.Address]*itokencontract.Itokencontract{},
		rpcInstancesCache:      map[int64]*ethclient.Client{},

		networker: cfg.NetworkConnector(),
	}
}

func (r *TokenContractReader) GetRPCInstance(chainID int64) (*ethclient.Client, error) {
	// Searching RPC instance in cache, if not found -- create new and store
	cacheInstance, ok := r.rpcInstancesCache[chainID]
	if ok {
		return cacheInstance, nil
	}

	// if specific chain is not cached yet -- getting network from connector
	// getting from connector moved here for reducing the great amount
	// of requests to network-svc
	// unlike FactoryReader, we can do requests to network-svc
	// only if chain was not found in cache; in FactoryReader we should always
	// pull all list of available networks

	network, err := r.networker.GetNetworkByChainID(chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get network by id", logan.F{
			"chain_id": chainID,
		})
	}
	if network == nil {
		return nil, errors.From(errors.New("network is nil"), logan.F{
			"chain_id": chainID,
		})
	}

	newInstance, err := ethclient.Dial(network.Data.Attributes.RpcUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert value into eth client", logan.F{
			"raw_url": network.Data.Attributes.RpcUrl,
		})
	}

	r.rpcInstancesCache[chainID] = newInstance
	return newInstance, nil

}

func (r *TokenContractReader) getContractInstance(address common.Address) (*itokencontract.Itokencontract, error) {
	// Searching contract instance in cache, if not found -- create new and store
	cacheInstance, ok := r.contractInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := itokencontract.NewItokencontract(*r.address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance for given address", logan.F{
			"address": address,
		})
	}

	r.contractInstancesCache[address] = newInstance
	return newInstance, nil
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

func (r *TokenContractReader) WithRPC(rpc *ethclient.Client) reader.TokenReader {
	r.rpc = rpc
	return r
}

func (r *TokenContractReader) validateParameters() error {
	//TODO: SHOULD WE VALIDATE `TO` PARAM?

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

func (r *TokenContractReader) GetUpdateEvents() ([]ethereum.UpdateEvent, error) {
	if err := r.validateParameters(); err != nil {
		return nil, err
	}

	events := make([]ethereum.UpdateEvent, 0)

	instance, err := r.getContractInstance(*r.address)
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
