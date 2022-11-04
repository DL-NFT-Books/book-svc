package helpers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/resources"
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

func SetMediaLinks(r *http.Request, banner, file *resources.Media) error {
	dconnector := DocumenterConnector(r)

	bannerLink, err := dconnector.GetDocumentLink(banner.Attributes.Key)
	if err != nil {
		return err
	}

	banner.Attributes.Url = &bannerLink.Data.Attributes.Url

	fileLink, err := dconnector.GetDocumentLink(file.Attributes.Key)
	if err != nil {
		return err
	}

	file.Attributes.Url = &fileLink.Data.Attributes.Url

	return nil
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
