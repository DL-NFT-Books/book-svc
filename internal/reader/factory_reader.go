package reader

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/ethereum"
)

type FactoryReader interface {
	From(from uint64) FactoryReader
	To(to uint64) FactoryReader
	WithAddress(address common.Address) FactoryReader

	GetDeployEvents() ([]ethereum.DeployEvent, error)
}
