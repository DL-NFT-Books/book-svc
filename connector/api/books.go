package api

import (
	"encoding/json"
	"fmt"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

const booksEndpoint = "books"

func (c *Connector) CreateBook(request resources.CreateBook) (*resources.CreateSignatureResponse, error) {
	var response resources.CreateSignatureResponse

	endpoint := fmt.Sprintf("%s/%s", c.baseUrl, booksEndpoint)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request")
	}

	if err = c.post(endpoint, requestAsBytes, &response); err != nil {
		return nil, errors.Wrap(err, "failed to create a book")
	}

	return &response, nil
}

func (c *Connector) UpdateBook(request resources.UpdateBook) error {
	endpoint := fmt.Sprintf("%s/%s/%s", c.baseUrl, booksEndpoint, request.ID)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}

func (c *Connector) ListBooks(request requests.GetBooksRequest) (*resources.BookListResponse, error) {
	var result resources.BookListResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s?%s", c.baseUrl, booksEndpoint, urlval.MustEncode(request))

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}

func (c *Connector) GetBookById(id int64) (*resources.BookResponse, error) {
	var result resources.BookResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, booksEndpoint, id)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": id})
	}

	return &result, nil
}
