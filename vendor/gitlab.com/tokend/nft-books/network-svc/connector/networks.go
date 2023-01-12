package connector

import (
	"fmt"
	"gitlab.com/tokend/nft-books/network-svc/connector/models"
	"gitlab.com/tokend/nft-books/network-svc/internal/data"
	"log"
)

const (
	networksEndpoint = "detailed"
)

func (c *Connector) GetNetworkByChainID(chainID int64) (*models.NetworkResponse, error) {
	var result data.Network

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%v", c.baseUrl, networksEndpoint, chainID)
	log.Println("CHAIN ID", chainID)
	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return result.ModelDetailed()
}
func (c *Connector) GetNetworks() (*models.NetworkListResponse, error) {
	var result []data.Network

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s", c.baseUrl, networksEndpoint)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	var networks models.NetworkListResponse

	for _, net := range result {
		modelNet, err := net.ModelDetailed()
		if err != nil {
			return nil, err
		}
		networks.Data = append(networks.Data, *modelNet)
	}

	return &networks, nil
}
