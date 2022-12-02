package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

type ListBooksRequest struct {
	pgdb.OffsetPageParams

	Status   *resources.DeployStatus `filter:"deploy_status"`
	Contract []string                `filter:"contract"`
	IDs      []int64                 `filter:"id"`
}

func NewListBooksRequest(r *http.Request) (ListBooksRequest, error) {
	var request ListBooksRequest
	err := urlval.Decode(r.URL.Query(), &request)
	return request, err
}
