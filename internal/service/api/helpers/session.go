package helpers

import (
	"errors"
	"net/http"
)

func ValidateJwt(r *http.Request) (string, error) {
	doorman := DoormanConnector(r)

	token, err := doorman.GetAuthToken(r)
	if err != nil {
		return "", errors.New("invalid token")
	}

	address, err := doorman.ValidateJwt(token)
	if err != nil {
		return "", errors.New("user does not have permission")
	}

	return address, nil
}
