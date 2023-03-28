package responses

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
	"github.com/dl-nft-books/book-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/book-svc/internal/service/api/requests"
	"github.com/dl-nft-books/book-svc/resources"
	"net/http"
)

// CreateLinks - return resources.Links structure with filled
// links from given url and pagination structure.
func CreateLinks(r *http.Request, request requests.ListBooksRequest) (*resources.Links, error) {
	count, err := helpers.GetBooksCount(r, &request)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get books count")
	}
	if count <= (request.OffsetPageParams.PageNumber+1)*request.OffsetPageParams.Limit {
		return &resources.Links{}, nil
	}
	return &resources.Links{
		Next: helpers.SetPageParams(*r.URL, pgdb.OffsetPageParams{
			PageNumber: request.OffsetPageParams.PageNumber + 1,
			Limit:      request.OffsetPageParams.Limit,
			Order:      request.OffsetPageParams.Order,
		}).String(),
	}, nil
}
