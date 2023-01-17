package responses

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/resources"
	"net/url"
)

// CreateLinks - return resources.Links structure with filled
// links from given url and pagination structure.
func CreateLinks(url *url.URL, params pgdb.OffsetPageParams) *resources.Links {
	links := &resources.Links{
		Next: helpers.SetPageParams(*url, pgdb.OffsetPageParams{
			PageNumber: params.PageNumber + 1,
			Limit:      params.Limit,
			Order:      params.Order,
		}).String(),
	}
	return links
}
