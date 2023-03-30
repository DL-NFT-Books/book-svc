package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/dl-nft-books/book-svc/resources"
)

const (
	MaxTitleLength       = 64
	MaxDescriptionLength = 500
)

type CreateBookRequest struct {
	Data resources.CreateBook `json:"data"`
}

func NewCreateBookRequest(r *http.Request) (CreateBookRequest, error) {
	var request CreateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return CreateBookRequest{}, errors.Wrap(err, "failed to decode request")
	}

	return request, request.validate()
}

func (r CreateBookRequest) validate() error {
	return validation.Errors{
		"/data/attributes/description": validation.Validate(
			&r.Data.Attributes.Description,
			validation.Required,
			validation.Length(1, MaxDescriptionLength)),
		"/included/banner/attributes/name":      validation.Validate(&r.Data.Attributes.Banner.Attributes.Name, validation.Required),
		"/included/banner/attributes/mime_type": validation.Validate(&r.Data.Attributes.Banner.Attributes.MimeType, validation.Required),
		"/included/banner/attributes/key":       validation.Validate(&r.Data.Attributes.Banner.Attributes.Key, validation.Required),

		"/included/file/attributes/name":      validation.Validate(&r.Data.Attributes.File.Attributes.Name, validation.Required),
		"/included/file/attributes/mime_type": validation.Validate(&r.Data.Attributes.File.Attributes.MimeType, validation.Required),
		"/included/file/attributes/key":       validation.Validate(&r.Data.Attributes.File.Attributes.Key, validation.Required),
		"/data/attributes/chain_ids": validation.Validate(
			&r.Data.Attributes.ChainIds,
			validation.Required,
			validation.Length(1, MaxDescriptionLength)),
	}.Filter()
}
