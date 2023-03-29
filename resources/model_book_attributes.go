/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type BookAttributes struct {
	Banner Media `json:"banner"`
	// Book creation time
	CreatedAt time.Time `json:"created_at"`
	// Book description
	Description string        `json:"description"`
	File        Media         `json:"file"`
	Networks    []BookNetwork `json:"networks"`
}
