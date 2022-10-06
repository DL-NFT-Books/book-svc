/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Book struct {
	Key
	Attributes BookAttributes `json:"attributes"`
}
type BookResponse struct {
	Data     Book     `json:"data"`
	Included Included `json:"included"`
}

type BookListResponse struct {
	Data     []Book   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustBook - returns Book from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBook(key Key) *Book {
	var book Book
	if c.tryFindEntry(key, &book) {
		return &book
	}
	return nil
}
