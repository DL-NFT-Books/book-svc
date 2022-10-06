/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BookAttributes struct {
	Banner Media `json:"banner"`
	// Book description
	Description string `json:"description"`
	File        Media  `json:"file"`
	// Book price ($)
	Price int32 `json:"price"`
	// Book title
	Title string `json:"title"`
}
