package connector

import (
	"encoding/json"
	"fmt"
	"github.com/dl-nft-books/book-svc/connector/models"
	"github.com/dl-nft-books/book-svc/internal/service/api/requests"
	"github.com/dl-nft-books/book-svc/resources"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

const booksEndpoint = "books"

func (c *Connector) UpdateBook(params models.UpdateBookParams) error {
	request := requests.UpdateBookRequest{
		ID: cast.ToInt64(params.Id),
		Data: resources.UpdateBook{
			Key: resources.NewKeyInt64(cast.ToInt64(params.Id), resources.BOOKS),
			Attributes: resources.UpdateBookAttributes{
				Banner:      params.Banner,
				Description: params.Description,
				File:        params.File,
			},
		},
	}

	endpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, booksEndpoint, request.ID)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}

func (c *Connector) ListBooks(request models.ListBooksParams) (*models.ListBooksResponse, error) {
	var result models.ListBooksResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s?%s", c.baseUrl, booksEndpoint, urlval.MustEncode(request))

	// getting response
	if _, err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}

func (c *Connector) GetBookById(bookId int64, chainId ...int64) (*models.GetBookResponse, error) {
	var result models.GetBookResponse

	var query struct {
		ChainId []int64 `filter:"chain_id"`
	}
	query.ChainId = chainId
	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d?%s", c.baseUrl, booksEndpoint, bookId, urlval.MustEncode(query))

	// getting response
	found, err := c.get(fullEndpoint, &result)
	if err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": bookId})
	}
	if !found {
		return nil, nil
	}

	return &result, nil
}
