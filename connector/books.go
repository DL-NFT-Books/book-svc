package connector

import (
	"encoding/json"
	"fmt"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/book-svc/connector/models"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

const booksEndpoint = "books"

func (c *Connector) CreateBook(params models.CreateBookParams) (createdId int64, err error) {
	request := requests.CreateBookRequest{
		Data: resources.CreateBook{
			Key: resources.NewKeyInt64(0, resources.BOOKS),
			Attributes: resources.CreateBookAttributes{
				Banner:      params.Banner,
				Description: params.Description,
				File:        params.File,
				Price:       params.Price,
				Title:       params.Title,
				TokenName:   params.TokenName,
				TokenSymbol: params.TokenSymbol,
			},
		},
	}

	var response resources.CreateBookResponse

	endpoint := fmt.Sprintf("%s/%s", c.baseUrl, booksEndpoint)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return 0, errors.Wrap(err, "failed to marshal request")
	}

	if err = c.post(endpoint, requestAsBytes, &response); err != nil {
		return 0, errors.Wrap(err, "failed to create a book")
	}

	createdBookId := cast.ToInt64(response.Data.ID)
	return createdBookId, nil
}

func (c *Connector) UpdateBook(params models.UpdateBookParams) error {
	request := requests.UpdateBookRequest{
		ID: params.Id,
		Data: resources.UpdateBook{
			Key: resources.NewKeyInt64(params.Id, resources.BOOKS),
			Attributes: resources.UpdateBookAttributes{
				Banner:          params.Banner,
				Description:     params.Description,
				File:            params.File,
				Title:           params.Title,
				ContractAddress: params.ContractAddress,
				DeployStatus:    params.DeployStatus,
				TokenSymbol:     params.Symbol,
				Price:           params.Price,
			},
		},
	}

	endpoint := fmt.Sprintf("%s/%s/%s", c.baseUrl, booksEndpoint, request.ID)
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

func (c *Connector) GetBookById(id int64) (*models.GetBookResponse, error) {
	var result models.GetBookResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, booksEndpoint, id)

	// getting response
	found, err := c.get(fullEndpoint, &result)
	if err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": id})
	}
	if !found {
		return nil, nil
	}

	return &result, nil
}
