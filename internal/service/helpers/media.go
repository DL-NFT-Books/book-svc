package helpers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/resources"
	"strings"
)

var (
	AllowedFileExtensions   = []string{"pdf"}
	AllowedBannerExtensions = []string{"png", "jpg", "jpeg"}
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

func CheckMediaTypes(banner, file string) error {
	bannerExt := strings.Split(banner, "/")
	if len(bannerExt) != 2 {
		return errors.New("invalid mime type")
	}

	err := checkBannerMimeType(bannerExt[1])
	if err != nil {
		return err
	}

	fileExt := strings.Split(file, "/")
	if len(bannerExt) != 2 {
		return errors.New("invalid mime type")
	}

	return checkFileMimeType(fileExt[1])
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
