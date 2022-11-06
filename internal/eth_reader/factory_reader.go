package eth_reader

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/solidity/generated/itokenfactory"
)

type FactoryContractReader struct {
	rpc *ethclient.Client
}

func NewFactoryContractReader(rpc *ethclient.Client) FactoryContractReader {
	return FactoryContractReader{
		rpc: rpc,
	}
}

type DeployEvent struct {
	Address         common.Address
	BlockNumber     uint64
	Name, Symbol    string
	TokenId, Status uint64
}

func (r *FactoryContractReader) GetDeployEvents(
	contract common.Address,
	startBlock,
	endBlock uint64,
) (
	events []DeployEvent,
	lastBlock uint64,
	err error,
) {
	instance, err := itokenfactory.NewItokenfactory(contract, r.rpc)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to initialize token factory instance")
	}

	iterator, err := instance.FilterTokenContractDeployed(
		&bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		},
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return nil, 0, NullIteratorErr
	}

	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event

		if event != nil {
			receipt, err := r.rpc.TransactionReceipt(context.Background(), event.Raw.TxHash)
			if err != nil {
				return nil, 0, errors.Wrap(err, "failed to get tx receipt", logan.F{
					"tx_hash": event.Raw.TxHash.String(),
				})
			}

			events = append(events, DeployEvent{
				Address:     event.NewTokenContractAddr,
				BlockNumber: event.Raw.BlockNumber,
				Name:        event.TokenName,
				Symbol:      event.TokenSymbol,
				TokenId:     event.TokenContractId.Uint64(),
				Status:      receipt.Status,
			})
		}

		lastBlock = event.Raw.BlockNumber
	}

	return
}
