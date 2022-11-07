package requests

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type GetBookByIDRequest struct {
	ID int64 `json:"id"`
}

func NewGetBookByIDRequest(r *http.Request) (GetBookByIDRequest, error) {
	var req GetBookByIDRequest

	id := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(id); err != nil {
		return req, errors.New("id is not an integer")
	}

	req.ID = cast.ToInt64(id)
	return req, nil
}
