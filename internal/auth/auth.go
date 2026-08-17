package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers *http.Header) (string, error) {
	val := headers.Get("X-API-Key")
	if val == "" {
		return "", errors.New("API key not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid API key format")
	}
	if vals[0] != "APIKey" {
		return "", errors.New("Invalid API key format")
	}
	return vals[1], nil
}