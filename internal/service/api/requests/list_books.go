package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type ListBooksRequest struct {
	pgdb.OffsetPageParams

	Contract []string `filter:"contract"`
	Id       []int64  `filter:"id"`
	ChainId  []int64  `filter:"chain_id"`
}

func NewListBooksRequest(r *http.Request) (ListBooksRequest, error) {
	var request ListBooksRequest
	err := urlval.Decode(r.URL.Query(), &request)
	return request, err
}
