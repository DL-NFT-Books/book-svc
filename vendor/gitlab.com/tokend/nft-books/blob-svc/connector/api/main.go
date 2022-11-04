package api

import (
	"encoding/json"
	"net/url"

	"gitlab.com/distributed_lab/json-api-connector/base"
	"gitlab.com/distributed_lab/json-api-connector/client"
)

const (
	DocumentEndpoint = "/integrations/documents"
)

type Connector struct {
	base   *base.Connector
	client client.Client
	token  string
}

func NewConnector(client client.Client, token string) *Connector {
	return &Connector{
		client: client,
		base:   base.NewConnector(client),
		token:  token,
	}
}

func (c *Connector) Get(endpoint *url.URL, dst interface{}) (err error) {
	response, err := c.base.Get(endpoint)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}
