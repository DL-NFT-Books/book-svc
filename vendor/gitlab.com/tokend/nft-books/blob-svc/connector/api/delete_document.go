package api

import (
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *Connector) DeleteDocument(key string) (int, error) {
	// forming endpoint
	endpoint := fmt.Sprintf("%s/%s/%s", c.baseUrl, DocumentEndpoint, key)

	// creating request
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "failed to build request")
	}

	// adding token for doorman auth
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// creating new client (!) and sending request
	resp, err := c.client.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return http.StatusBadRequest, err
	}

	return resp.StatusCode, nil
}
