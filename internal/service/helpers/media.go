package helpers

import (
	"encoding/json"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

func MarshalMedia(media ...resources.Media) []string {
	var res []string

	for _, v := range media {
		raw, err := json.Marshal(v)
		if err != nil {
			return nil
		}

		res = append(res, string(raw))
	}

	return res
}

func UnmarshalMedia(media string) (resources.Media, error) {
	var res resources.Media
	err := json.Unmarshal([]byte(media), &res)
	return res, err
}
