package helpers

import (
	"net/url"
	"strconv"

	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	pageLimitParam  = "page[limit]"
	pageNumberParam = "page[number]"
	pageOrderParam  = "page[order]"
)

func SetPageParams(u url.URL, params pgdb.OffsetPageParams) *url.URL {
	query := u.Query()
	if params.Limit != 0 {
		query.Set(pageLimitParam, strconv.Itoa(int(params.Limit)))
	}
	if params.PageNumber != 0 {
		query.Set(pageNumberParam, strconv.Itoa(int(params.PageNumber)))
	} else {
		query.Del(pageNumberParam)
	}
	if len(params.Order) != 0 {
		query.Set(pageOrderParam, params.Order)
	}
	u.RawQuery = query.Encode()
	return &u
}
