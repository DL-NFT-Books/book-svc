package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/resources"
	"net/http"
)

const (
	S3KeyLength          = 36
	MinExtLength         = 3
	MaxExtLength         = 4
	MaxTitleLength       = 64
	MaxDescriptionLength = 500
)

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
		"/data/attributes/title":       validation.Validate(&r.Data.Attributes.Title, validation.Required, validation.Length(1, MaxTitleLength)),
		"/data/attributes/description": validation.Validate(&r.Data.Attributes.Description, validation.Required, validation.Length(1, MaxDescriptionLength)),
		"/data/attributes/price":       validation.Validate(&r.Data.Attributes.Price, validation.Required),

		"/included/banner/attributes/name":      validation.Validate(&r.Banner.Attributes.Name, validation.Required),
		"/included/banner/attributes/mime_type": validation.Validate(&r.Banner.Attributes.MimeType, validation.Required),
		"included/banner/attributes/key": validation.Validate(
			&r.Banner.Attributes.Key,
			validation.Required,
			validation.Length(S3KeyLength+1+MinExtLength, S3KeyLength+1+MaxExtLength)), //include '.'

		"/included/file/attributes/name":      validation.Validate(&r.File.Attributes.Name, validation.Required),
		"/included/file/attributes/mime_type": validation.Validate(&r.File.Attributes.MimeType, validation.Required),
		"/included/file/attributes/key": validation.Validate(
			&r.File.Attributes.Key,
			validation.Required,
			validation.Length(S3KeyLength+1+MinExtLength, S3KeyLength+1+MaxExtLength)), //include '.'
	}.Filter()
}
