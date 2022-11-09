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
	var req UpdateBookRequest

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return req, errors.New("id is not an integer")
	}
	req.ID = int64(id)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.Wrap(err, "failed to decode request")
	}

	return req, nil
}
