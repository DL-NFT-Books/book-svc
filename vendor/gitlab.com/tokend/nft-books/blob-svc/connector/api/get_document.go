package api

import (
	"fmt"
	"net/url"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/blob-svc/connector/models"
)

func (c *Connector) GetDocumentLink(key string) (models.LinkResponse, error) {
	var response models.LinkResponse

	parsedUrl, err := url.Parse(fmt.Sprintf("%s/%s", DocumentEndpoint, key))
	if err != nil {
		return response, errors.Wrap(err, "failed to parse document url")
	}

	err = c.Get(parsedUrl, &response)
	if err != nil {
		return response, errors.Wrap(err, "failed to get document")
	}

	return response, nil
}
