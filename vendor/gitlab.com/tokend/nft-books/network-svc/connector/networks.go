package connector

import (
	"fmt"
	"gitlab.com/tokend/nft-books/network-svc/connector/models"
)

const (
	networksEndpoint = "detailed"
)

func (c *Connector) GetNetworkByChainID(chainID int64) (*models.NetworkResponse, error) {
	var result models.NetworkResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%v", c.baseUrl, networksEndpoint, chainID)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}
func (c *Connector) GetNetworks() (*models.NetworkListResponse, error) {
	var result models.NetworkListResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s", c.baseUrl, networksEndpoint)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}
