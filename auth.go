package auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

// ErrUnauthorized means that the requester is not authorized
var ErrUnauthorized = errors.New("unauthorized")

// Authorizer is the interface that wraps the basic Authrorize method.
type Authorizer interface {
	Authorize(r *http.Request) error
}

// BasicAuthorizer authorize using Basic HTTP authentication scheme
type BasicAuthorizer struct {
	username string
	password string
}

// NewBasicAuthorizer returns a BasicAuthorizer with specific username and password
func NewBasicAuthorizer(username string, password string) *BasicAuthorizer {
	return &BasicAuthorizer{username: username, password: password}
}

// Authorize returns error if authorization header is not match with authorizer credentials
func (a *BasicAuthorizer) Authorize(r *http.Request) error {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return ErrUnauthorized
	}

	decoded, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		return err
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 || a.username != credentials[0] || a.password != credentials[1] {
		return ErrUnauthorized
	}

	return nil
}
