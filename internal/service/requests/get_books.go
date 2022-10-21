package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type GetBooksRequest struct {
	pgdb.OffsetPageParams
}

func NewGetBooksRequest(r *http.Request) (GetBooksRequest, error) {
	var req GetBooksRequest
	err := urlval.Decode(r.URL.Query(), &req)
	return req, err
}
