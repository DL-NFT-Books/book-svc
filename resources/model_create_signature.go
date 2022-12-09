/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateSignature struct {
	Key
	Attributes    CreateSignatureAttributes    `json:"attributes"`
	Relationships CreateSignatureRelationships `json:"relationships"`
}
type CreateSignatureResponse struct {
	Data     CreateSignature `json:"data"`
	Included Included        `json:"included"`
}

type CreateSignatureListResponse struct {
	Data     []CreateSignature `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustCreateSignature - returns CreateSignature from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateSignature(key Key) *CreateSignature {
	var createSignature CreateSignature
	if c.tryFindEntry(key, &createSignature) {
		return &createSignature
	}
	return nil
}
