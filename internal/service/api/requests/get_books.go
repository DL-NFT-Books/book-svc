package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

type GetBooksRequest struct {
	pgdb.OffsetPageParams

	Status  *resources.DeployStatus `filter:"deploy_status"`
	ChainID *int64                  `filter:"chain_id"`
	IDs     []int64                 `filter:"id"`
}

func NewGetBooksRequest(r *http.Request) (GetBooksRequest, error) {
	var req GetBooksRequest
	err := urlval.Decode(r.URL.Query(), &req)
	return req, err
}
