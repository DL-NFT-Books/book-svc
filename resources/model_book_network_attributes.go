/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BookNetworkAttributes struct {
	// Networks chain id
	ChainId int64 `json:"chain_id"`
	// Token contract address
	ContractAddress string `json:"contract_address"`
	// status of a book deployment
	DeployStatus DeployStatus `json:"deploy_status"`
	// id from the contract that corresponds to the given book
	TokenId int64 `json:"token_id"`
}
