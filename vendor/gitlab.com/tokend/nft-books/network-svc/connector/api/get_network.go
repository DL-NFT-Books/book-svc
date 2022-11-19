package api

import (
	"fmt"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/network-svc/connector/models"
)

func (c *Connector) GetNetworkByChainID(chainID int64) (*models.NetworkResponse, error) {
	var result models.NetworkResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%v", c.baseUrl, NetworksEndpoint, chainID)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"chain_id": chainID})
	}

	return &result, nil
}
