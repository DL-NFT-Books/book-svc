package requests

import (
	"encoding/json"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/tokend/nft-books/book-svc/resources"
)

const (
	MaxTitleLength       = 64
	MaxDescriptionLength = 500
)

var AddressRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]{40}$")

type CreateBookRequest struct {
	Data     resources.Book     `json:"data"`
	Included resources.Included `json:"included"`
	File     *resources.Media   `json:"file"`
	Banner   *resources.Media   `json:"banner"`
}

func NewCreateBookRequest(r *http.Request) (CreateBookRequest, error) {
	var req CreateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return CreateBookRequest{}, errors.Wrap(err, "failed to decode request")
	}

	req.File = req.Included.MustMedia(req.Data.Relationships.File.Data.GetKey())
	req.Banner = req.Included.MustMedia(req.Data.Relationships.Banner.Data.GetKey())

	return req, req.validate()
}

func (r CreateBookRequest) validate() error {
	return validation.Errors{
		"/data/attributes/title": validation.Validate(
			&r.Data.Attributes.Title,
			validation.Required,
			validation.Length(1, MaxTitleLength)),
		"/data/attributes/description": validation.Validate(
			&r.Data.Attributes.Description,
			validation.Required,
			validation.Length(1, MaxDescriptionLength)),
		"/data/attributes/contract_address": validation.Validate(
			&r.Data.Attributes.ContractAddress,
			validation.Required,
			validation.Match(AddressRegexp)),
		"/data/attributes/contract_name": validation.Validate(&r.Data.Attributes.ContractName, validation.Required),

		"/included/banner/attributes/name":      validation.Validate(&r.Banner.Attributes.Name, validation.Required),
		"/included/banner/attributes/mime_type": validation.Validate(&r.Banner.Attributes.MimeType, validation.Required),
		"/included/banner/attributes/key":       validation.Validate(&r.Banner.Attributes.Key, validation.Required),

		"/included/file/attributes/name":      validation.Validate(&r.File.Attributes.Name, validation.Required),
		"/included/file/attributes/mime_type": validation.Validate(&r.File.Attributes.MimeType, validation.Required),
		"/included/file/attributes/key":       validation.Validate(&r.File.Attributes.Key, validation.Required),
	}.Filter()
}
