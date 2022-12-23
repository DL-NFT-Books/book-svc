/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type BookAttributes struct {
	Banner  Media `json:"banner"`
	ChainId int64 `json:"chain_id"`
	// Token contract address
	ContractAddress string `json:"contract_address"`
	// Token contract name
	ContractName string `json:"contract_name"`
	// Token contract symbol
	ContractSymbol string `json:"contract_symbol"`
	// Token contract version
	ContractVersion string `json:"contract_version"`
	// Book creation time
	CreatedAt time.Time `json:"created_at"`
	// status of a book deployment
	DeployStatus DeployStatus `json:"deploy_status"`
	// Book description
	Description string `json:"description"`
	File        Media  `json:"file"`
	// Book price ($)
	Price string `json:"price"`
	// Book title
	Title string `json:"title"`
	// id from the contract that corresponds to the given book
	TokenId int64 `json:"token_id"`
}
