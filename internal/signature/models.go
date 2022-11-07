package signature

import "math/big"

type EIP712DomainData struct {
	ContractName     string
	ContractVersion  string
	VerifyingAddress string
	ChainID          int64
}

type CreateInfo struct {
	TokenContractId  int64
	TokenName        string
	TokenSymbol      string
	PricePerOneToken *big.Int

	HashedTokenName   []byte
	HashedTokenSymbol []byte
}
