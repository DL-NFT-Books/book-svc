package api

import (
	"fmt"
	"net/http"
	"net/url"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *Connector) DeleteDocument(key string) (int, error) {
	// forming endpoint
	parsedUrl, err := url.Parse(fmt.Sprintf("%s/%s", DocumentEndpoint, key))
	if err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "failed to parse document url")
	}

	fullEndpoint, err := c.client.Resolve(parsedUrl)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// creating request
	req, err := http.NewRequest(http.MethodDelete, fullEndpoint, nil)
	if err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "failed to build request")
	}

	// adding token for doorman auth
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// creating new client (!) and sending request
	newCli := http.Client{}
	resp, err := newCli.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return http.StatusBadRequest, err
	}

	return resp.StatusCode, nil
}
