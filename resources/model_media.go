/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Media struct {
	Key
	Attributes MediaAttributes `json:"attributes"`
}
type MediaResponse struct {
	Data     Media    `json:"data"`
	Included Included `json:"included"`
}

type MediaListResponse struct {
	Data     []Media  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustMedia - returns Media from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustMedia(key Key) *Media {
	var media Media
	if c.tryFindEntry(key, &media) {
		return &media
	}
	return nil
}
