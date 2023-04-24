package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	DocumentEndpoint = "documents"
)

type Connector struct {
	baseUrl string
	client  *http.Client
	token   string
}

func NewConnector(url, token string) *Connector {
	return &Connector{
		client:  http.DefaultClient,
		baseUrl: url,
		token:   token,
	}
}

func (c *Connector) Get(endpoint string, dst interface{}) (err error) {
	// creating request
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create connector request")
	}

	// sending request
	response, err := c.client.Do(request)
	if err != nil {
		return errors.Wrap(err, "failed to process request")
	}

	defer func(Body io.ReadCloser) {
		if tempErr := Body.Close(); tempErr != nil {
			err = tempErr
		}
	}(response.Body)

	// parsing response
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	return json.Unmarshal(raw, dst)
}
