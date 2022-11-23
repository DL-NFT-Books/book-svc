package reader

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/ethereum"
)

type FactoryReader interface {
	From(from uint64) FactoryReader
	To(to uint64) FactoryReader
	WithAddress(address common.Address) FactoryReader
	WithRPC(rpc *ethclient.Client) FactoryReader

	GetRPCInstance(rawURL string) (*ethclient.Client, error)
	GetDeployEvents() ([]ethereum.DeployEvent, error)
}
