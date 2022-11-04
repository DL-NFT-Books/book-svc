/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateBook struct {
	Key
	Attributes CreateBookAttributes `json:"attributes"`
}
type CreateBookResponse struct {
	Data     CreateBook `json:"data"`
	Included Included   `json:"included"`
}

type CreateBookListResponse struct {
	Data     []CreateBook `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustCreateBook - returns CreateBook from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateBook(key Key) *CreateBook {
	var createBook CreateBook
	if c.tryFindEntry(key, &createBook) {
		return &createBook
	}
	return nil
}
