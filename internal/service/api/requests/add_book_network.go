package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"

	"github.com/dl-nft-books/book-svc/resources"
)

type AddBookNetworkRequest struct {
	Data   []resources.BookNetwork `json:"data"`
	BookId int64                   `json:"-"`
}

func NewAddBookNetworkRequest(r *http.Request) (AddBookNetworkRequest, error) {
	var request AddBookNetworkRequest
	id := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(id); err != nil {
		return request, errors.New("id is not an integer")
	}
	request.BookId = cast.ToInt64(id)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return AddBookNetworkRequest{}, errors.Wrap(err, "failed to decode request")
	}

	return request, request.validate()
}

func (r AddBookNetworkRequest) validate() error {
	return validation.Errors{"/data/attributes/networks": validation.Validate(
		&r.Data,
		validation.Required,
		validation.Length(1, MaxDescriptionLength)),
	}.Filter()
}
