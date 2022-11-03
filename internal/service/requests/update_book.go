package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

type UpdateBookRequest struct {
	ID       int64                `json:"id"`
	Data     resources.UpdateBook `json:"data"`
	Included resources.Included   `json:"included"`
	File     *resources.Media     `json:"file"`
	Banner   *resources.Media     `json:"banner"`
}

func NewUpdateBookRequest(r *http.Request) (UpdateBookRequest, error) {
	var req UpdateBookRequest

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return req, errors.New("id is not an integer")
	}
	req.ID = int64(id)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.Wrap(err, "failed to decode request")
	}

	req.File = req.Included.MustMedia(req.Data.Relationships.File.Data.GetKey())
	req.Banner = req.Included.MustMedia(req.Data.Relationships.Banner.Data.GetKey())

	return req, req.validate()
}

func (r UpdateBookRequest) validate() error {
	return validation.Errors{
		"/data/attributes/title": validation.Validate(
			&r.Data.Attributes.Title,
			validation.Required,
			validation.Length(1, MaxTitleLength)),
		"/data/attributes/description": validation.Validate(
			&r.Data.Attributes.Description,
			validation.Required,
			validation.Length(1, MaxDescriptionLength)),

		"/included/banner/attributes/name":      validation.Validate(&r.Banner.Attributes.Name, validation.Required),
		"/included/banner/attributes/mime_type": validation.Validate(&r.Banner.Attributes.MimeType, validation.Required),
		"/included/banner/attributes/key":       validation.Validate(&r.Banner.Attributes.Key, validation.Required),

		"/included/file/attributes/name":      validation.Validate(&r.File.Attributes.Name, validation.Required),
		"/included/file/attributes/mime_type": validation.Validate(&r.File.Attributes.MimeType, validation.Required),
		"/included/file/attributes/key":       validation.Validate(&r.File.Attributes.Key, validation.Required),
	}.Filter()
}
