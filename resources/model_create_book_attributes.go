/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateBookAttributes struct {
	Banner Media `json:"banner"`
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
}
