package reader

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/ethereum"
)

type TokenReader interface {
	From(from uint64) TokenReader
	To(to uint64) TokenReader
	WithAddress(address common.Address) TokenReader
	WithRPC(rpc *ethclient.Client) TokenReader

	GetRPCInstance(chainID int64) (*ethclient.Client, error)
	GetUpdateEvents() ([]ethereum.UpdateEvent, error)
}
