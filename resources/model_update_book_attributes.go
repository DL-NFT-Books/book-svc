/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateBookAttributes struct {
	Banner *Media `json:"banner,omitempty"`
	// Book description
	Description *string `json:"description,omitempty"`
	File        *Media  `json:"file,omitempty"`
	// Book title
	Title           *string       `json:"title,omitempty"`
	ContractAddress *string       `json:"contract_address,omitempty"`
	DeployStatus    *DeployStatus `json:"deploy_status,omitempty"`
	// Price per one token ($)
	Price       *string `json:"price"`
	TokenSymbol *string `json:"token_symbol"`
}
