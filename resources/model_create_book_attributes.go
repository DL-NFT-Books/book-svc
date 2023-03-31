/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateBookAttributes struct {
	Banner Media `json:"banner"`
	// Book description
	Description string        `json:"description"`
	File        Media         `json:"file"`
	Networks    []BookNetwork `json:"networks"`
}
