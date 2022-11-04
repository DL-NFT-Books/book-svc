/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateBookAttributes struct {
	Banner Media `json:"banner"`
	// Token contract address
	ContractAddress string `json:"contract_address"`
	// Book description
	Description string `json:"description"`
	File        Media  `json:"file"`
	// Book title
	Title string `json:"title"`
}
