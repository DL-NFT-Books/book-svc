/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type KeyResponse struct {
	Key
	Attributes KeyResponseAttributes `json:"attributes"`
}
type KeyResponseResponse struct {
	Data     KeyResponse `json:"data"`
	Included Included    `json:"included"`
}

type KeyResponseListResponse struct {
	Data     []KeyResponse `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustKeyResponse - returns KeyResponse from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustKeyResponse(key Key) *KeyResponse {
	var keyResponse KeyResponse
	if c.tryFindEntry(key, &keyResponse) {
		return &keyResponse
	}
	return nil
}
