package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/resources"
	"net/http"
)

const AllowedS3KeyLength = 36

type CreateBookRequest struct {
	Data resources.Book `json:"data"`
}

func NewCreateBookRequest(r *http.Request) (resources.Book, error) {
	var req CreateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return resources.Book{}, errors.Wrap(err, "failed to decode request")
	}

	return req.Data, req.validate()
}

func (r CreateBookRequest) validate() error {
	return validation.Errors{
		"/data/attributes/title":       validation.Validate(&r.Data.Attributes.Title, validation.Required),
		"/data/attributes/description": validation.Validate(&r.Data.Attributes.Description, validation.Required),
		"/data/attributes/price":       validation.Validate(&r.Data.Attributes.Price, validation.Required, validation.Min(0)),

		"/data/relationships/banner/attributes/name":      validation.Validate(&r.Data.Relationships.Banner.Attributes.Name, validation.Required),
		"/data/relationships/banner/attributes/mime_type": validation.Validate(&r.Data.Relationships.Banner.Attributes.MimeType, validation.Required),
		"/data/relationships/banner/attributes/key":       validation.Validate(&r.Data.Relationships.Banner.Attributes.Key, validation.Required, validation.Length(AllowedS3KeyLength, AllowedS3KeyLength)),

		"/data/relationships/file/attributes/name":      validation.Validate(&r.Data.Relationships.File.Attributes.Name, validation.Required),
		"/data/relationships/file/attributes/mime_type": validation.Validate(&r.Data.Relationships.File.Attributes.MimeType, validation.Required),
		"/data/relationships/file/attributes/key":       validation.Validate(&r.Data.Relationships.File.Attributes.Key, validation.Required, validation.Length(AllowedS3KeyLength, AllowedS3KeyLength)),
	}.Filter()
}
