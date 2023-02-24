package responses

import (
	"fmt"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"
	"net/http"
)

// CreateLinks - return resources.Links structure with filled
// links from given url and pagination structure.
func CreateLinks(r *http.Request, request requests.ListBooksRequest) (*resources.Links, error) {
	count, err := helpers.GetBooksCount(r, &request)
	fmt.Println(count)
	if err != nil {
		return nil, err
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
