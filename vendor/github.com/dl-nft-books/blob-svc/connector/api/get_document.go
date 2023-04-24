package api

import (
	"fmt"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/dl-nft-books/blob-svc/connector/models"
)

func (c *Connector) GetDocumentLink(key string) (models.LinkResponse, error) {
	var response models.LinkResponse

	err := c.Get(fmt.Sprintf("%s/%s/%s", c.baseUrl, DocumentEndpoint, key), &response)
	if err != nil {
		return response, errors.Wrap(err, "failed to get document")
	}

	return response, nil
}
