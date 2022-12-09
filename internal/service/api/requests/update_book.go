package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

type UpdateBookRequest struct {
	ID   int64                `json:"id"`
	Data resources.UpdateBook `json:"data"`
}

func NewUpdateBookRequest(r *http.Request) (UpdateBookRequest, error) {
	var request UpdateBookRequest

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return request, errors.New("id is not an integer")
	}
	request.ID = int64(id)

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to decode request")
	}

	return request, nil
}
