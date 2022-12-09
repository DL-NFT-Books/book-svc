/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateBook struct {
	Key
	Attributes UpdateBookAttributes `json:"attributes"`
}
type UpdateBookResponse struct {
	Data     UpdateBook `json:"data"`
	Included Included   `json:"included"`
}

type UpdateBookListResponse struct {
	Data     []UpdateBook `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustUpdateBook - returns UpdateBook from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateBook(key Key) *UpdateBook {
	var updateBook UpdateBook
	if c.tryFindEntry(key, &updateBook) {
		return &updateBook
	}
	return nil
}
