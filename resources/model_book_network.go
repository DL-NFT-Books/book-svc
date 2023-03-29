/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BookNetwork struct {
	Key
	Attributes BookNetworkAttributes `json:"attributes"`
}
type BookNetworkResponse struct {
	Data     BookNetwork `json:"data"`
	Included Included    `json:"included"`
}

type BookNetworkListResponse struct {
	Data     []BookNetwork `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustBookNetwork - returns BookNetwork from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBookNetwork(key Key) *BookNetwork {
	var bookNetwork BookNetwork
	if c.tryFindEntry(key, &bookNetwork) {
		return &bookNetwork
	}
	return nil
}
