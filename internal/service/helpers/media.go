package helpers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/resources"
	"net/http"
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

func CheckMediaTypes(r *http.Request, bannerExt, fileExt string) error {
	err := checkBannerMimeType(bannerExt, r)
	if err != nil {
		return err
	}

	return checkFileMimeType(fileExt, r)
}

func checkBannerMimeType(ext string, r *http.Request) error {
	for _, el := range MimeTypes(r).AllowedBannerMimeTypes {
		if el == ext {
			return nil
		}
	}
	return errors.New("invalid banner extension")
}

func checkFileMimeType(ext string, r *http.Request) error {
	for _, el := range MimeTypes(r).AllowedFileMimeTypes {
		if el == ext {
			return nil
		}
	}
	return errors.New("invalid file extension")
}
