package connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Connector struct {
	token   string
	baseUrl string
	client  *http.Client
}

func NewConnector(authToken, serviceUrl string) *Connector {
	return &Connector{
		token:   authToken,
		baseUrl: serviceUrl,
		client:  http.DefaultClient,
	}
}

func (c *Connector) get(endpoint string, dst interface{}) (found bool, err error) {
	// creating request
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return false, errors.Wrap(err, "failed to create connector request")
	}

	// setting headers
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// sending request
	response, err := c.client.Do(request)
	if err != nil {
		return false, errors.Wrap(err, "failed to process request")
	}

	if response.StatusCode == http.StatusNotFound {
		return false, nil
	}

	defer func(Body io.ReadCloser) {
		if tempErr := Body.Close(); tempErr != nil {
			err = tempErr
		}
	}(response.Body)

	// parsing response
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, errors.Wrap(err, "failed to read response body")
	}

	return json.Unmarshal(raw, &dst) == nil, err
}

func (c *Connector) update(endpoint string, data []byte, dst interface{}) error {
	return c.upsert(http.MethodPatch, endpoint, data, dst)
}

func (c *Connector) post(endpoint string, data []byte, dst interface{}) error {
	return c.upsert(http.MethodPost, endpoint, data, dst)
}

// upsert is a function of POST or PATCH depending on what you choose as a method
func (c *Connector) upsert(method, endpoint string, data []byte, dst interface{}) error {
	// creating request
	bodyReader := bytes.NewReader(data)
	request, err := http.NewRequest(method, endpoint, bodyReader)
	if err != nil {
		return errors.Wrap(err, "failed to create connector request")
	}

	// setting headers
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// sending request
	response, err := c.client.Do(request)
	if err != nil {
		return errors.Wrap(err, "failed to process request")
	}
	if response == nil {
		return errors.New("failed to process request: response is nil")
	}

	defer func(Body io.ReadCloser) {
		if tempErr := Body.Close(); tempErr != nil {
			err = tempErr
		}
	}(response.Body)

	// If destination is nil, we do not read the response
	if dst == nil {
		return nil
	}

	// parsing response
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	return json.Unmarshal(raw, &dst)
}
