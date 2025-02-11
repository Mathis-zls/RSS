package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {insert Api key}
func GetAPIkey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first Part")
	}
	return vals[1], nil
}
