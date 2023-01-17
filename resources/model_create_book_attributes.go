/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateBookAttributes struct {
	Banner  Media `json:"banner"`
	ChainId int64 `json:"chain_id"`
	// Book description
	Description string `json:"description"`
	File        Media  `json:"file"`
	// Price per one token ($)
	Price string `json:"price"`
	// Book title
	Title string `json:"title"`
	// Token name
	TokenName string `json:"token_name"`
	// Token symbol
	TokenSymbol string `json:"token_symbol"`
	// Voucher token contract address, that can be used to claim free book
	VoucherToken *string `json:"voucher_token,omitempty"`
	// How many voucher tokens user have to pay that book
	VoucherTokenAmount *string `json:"voucher_token_amount,omitempty"`
}
