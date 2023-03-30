/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateBookAttributes struct {
	Banner   Media   `json:"banner"`
	ChainIds []int64 `json:"chain_ids"`
	// Book description
	Description string `json:"description"`
	File        Media  `json:"file"`
}
