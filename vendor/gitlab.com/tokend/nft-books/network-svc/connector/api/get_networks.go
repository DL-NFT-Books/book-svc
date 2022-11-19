package api

import (
	"fmt"

	"gitlab.com/tokend/nft-books/network-svc/connector/models"
)

func (c *Connector) GetNetworks() (*models.NetworkListResponse, error) {
	var result models.NetworkListResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s", c.baseUrl, NetworksEndpoint)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}
