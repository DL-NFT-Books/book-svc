package models

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type NetworkResponse struct {
	ChainId        int64             `json:"chain_id"`
	FactoryAddress string            `json:"factory_address"`
	FactoryName    string            `json:"factory_name"`
	FactoryVersion string            `json:"factory_version"`
	FirstBlock     int64             `json:"first_block"`
	Name           string            `json:"name"`
	RpcUrl         *ethclient.Client `json:"rpc_url"`
	TokenName      string            `json:"token_name"`
	TokenSymbol    string            `json:"token_symbol"`
	WsUrl          *ethclient.Client `json:"ws_url"`
}

type NetworkListResponse []NetworkResponse
