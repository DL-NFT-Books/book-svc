package ethereum

import "github.com/ethereum/go-ethereum/common"

type DeployEvent struct {
	Address         common.Address
	BlockNumber     uint64
	Name, Symbol    string
	TokenId, Status uint64
}
