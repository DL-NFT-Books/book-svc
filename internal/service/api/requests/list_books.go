package requests

import (
	"net/http"

	"github.com/dl-nft-books/book-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type ListBooksRequest struct {
	pgdb.OffsetPageParams

	Status   []resources.DeployStatus `filter:"deploy_status"`
	Contract []string                 `filter:"contract"`
	Id       []int64                  `filter:"id"`
	TokenId  []int64                  `filter:"token_id"`
	ChainId  []int64                  `filter:"chain_id"`
}

func NewListBooksRequest(r *http.Request) (ListBooksRequest, error) {
	var request ListBooksRequest
	err := urlval.Decode(r.URL.Query(), &request)
	return request, err
}
