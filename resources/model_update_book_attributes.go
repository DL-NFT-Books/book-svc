/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateBookAttributes struct {
	Banner *Media `json:"banner,omitempty"`
	// Address of a contract corresponding to this book
	ContractAddress *string `json:"contract_address,omitempty"`
	// Contract name (in most cases coincides with a title field)
	ContractName *string `json:"contract_name,omitempty"`
	// status of a book deployment
	DeployStatus *DeployStatus `json:"deploy_status,omitempty"`
	// Book description
	Description *string `json:"description,omitempty"`
	File        *Media  `json:"file,omitempty"`
	// Book floor price ($)
	FloorPrice *string `json:"floor_price,omitempty"`
	// Price per one token ($)
	Price *string `json:"price,omitempty"`
	// Book title
	Title *string `json:"title,omitempty"`
	// Token symbol
	TokenSymbol *string `json:"token_symbol,omitempty"`
	// Voucher token contract address, that can be used to claim free book
	VoucherToken *string `json:"voucher_token,omitempty"`
	// How many voucher tokens user have to pay that book
	VoucherTokenAmount *string `json:"voucher_token_amount,omitempty"`
}
