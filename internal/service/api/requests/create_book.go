package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/resources"
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
		"/data/attributes/title": validation.Validate(
			&r.Data.Attributes.Title,
			validation.Required,
			validation.Length(1, MaxTitleLength)),
		"/data/attributes/description": validation.Validate(
			&r.Data.Attributes.Description,
			validation.Required,
			validation.Length(1, MaxDescriptionLength)),
		"/data/attributes/token_name":   validation.Validate(&r.Data.Attributes.TokenName, validation.Required),
		"/data/attributes/token_symbol": validation.Validate(&r.Data.Attributes.TokenSymbol, validation.Required),
		"/data/attributes/price":        validation.Validate(&r.Data.Attributes.Price, validation.Required),

		"/included/banner/attributes/name":      validation.Validate(&r.Data.Attributes.Banner.Attributes.Name, validation.Required),
		"/included/banner/attributes/mime_type": validation.Validate(&r.Data.Attributes.Banner.Attributes.MimeType, validation.Required),
		"/included/banner/attributes/key":       validation.Validate(&r.Data.Attributes.Banner.Attributes.Key, validation.Required),

		"/included/file/attributes/name":      validation.Validate(&r.Data.Attributes.File.Attributes.Name, validation.Required),
		"/included/file/attributes/mime_type": validation.Validate(&r.Data.Attributes.File.Attributes.MimeType, validation.Required),
		"/included/file/attributes/key":       validation.Validate(&r.Data.Attributes.File.Attributes.Key, validation.Required),
	}.Filter()
}
