package helpers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

var (
	AllowedFileExtensions   = []string{"application/pdf"}
	AllowedBannerExtensions = []string{"img/png", "img/jpg", "img/jpeg"}
)

func MarshalMedia(media ...*resources.Media) []string {
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

func UnmarshalMedia(media ...string) ([]resources.Media, error) {
	var res []resources.Media
	var unmarshalledMedia resources.Media

	for _, value := range media {
		err := json.Unmarshal([]byte(value), &unmarshalledMedia)
		if err != nil {
			return nil, err
		}

		res = append(res, unmarshalledMedia)
	}
	return res, nil
}

func CheckMediaTypes(bannerExt, fileExt string) error {
	err := checkBannerMimeType(bannerExt)
	if err != nil {
		return err
	}

	return checkFileMimeType(fileExt)
}

func checkBannerMimeType(ext string) error {
	for _, el := range AllowedBannerExtensions {
		if el == ext {
			return nil
		}
	}
	return errors.New("invalid banner extension")
}

func checkFileMimeType(ext string) error {
	for _, el := range AllowedFileExtensions {
		if el == ext {
			return nil
		}
	}
	return errors.New("invalid file extension")
}
