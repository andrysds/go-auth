package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuthorizer(t *testing.T) {
	cases := []struct {
		authorization  string
		returningError bool
	}{
		{
			authorization:  "admin:admin",
			returningError: true,
		}, {
			authorization:  "admin admin",
			returningError: true,
		}, {
			authorization:  "Basic XXXYYYZZZ",
			returningError: true,
		}, {
			authorization:  "Basic cm9vdDpyb290",
			returningError: true,
		}, {
			authorization:  "Basic YWRtaW46YWRtaW4=",
			returningError: false,
		},
	}
	for _, c := range cases {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", c.authorization)

		authorizer := NewBasicAuthorizer("admin", "admin")
		err := authorizer.Authorize(req)

		if c.returningError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
