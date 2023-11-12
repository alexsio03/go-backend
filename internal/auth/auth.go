package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Example: Authorization : ApiKey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	str := headers.Get("Authorization")
	if str == "" {
		return "", errors.New("no authorization header")
	}

	vals := strings.Split(str, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization header given")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid first part of header")
	}
	return vals[1], nil
}
