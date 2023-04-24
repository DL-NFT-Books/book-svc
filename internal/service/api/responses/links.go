package responses

import (
	"github.com/dl-nft-books/book-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/book-svc/internal/service/api/requests"
	"github.com/dl-nft-books/book-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
	"net/http"
)

// CreateLinks - return resources.Links structure with filled
// links from given url and pagination structure.
func CreateLinks(r *http.Request, request requests.ListBooksRequest) (*resources.Links, error) {
	return &resources.Links{
		Next: helpers.SetPageParams(*r.URL, pgdb.OffsetPageParams{
			PageNumber: request.OffsetPageParams.PageNumber + 1,
			Limit:      request.OffsetPageParams.Limit,
			Order:      request.OffsetPageParams.Order,
		}).String(),
	}, nil
}
