package json_body

import (
	"github.com/stretchr/testify/assert"
	gfun "go-shopping/go/func"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
	httpTest "go-shopping/tests/lambler/http"
	"net/http"
	"testing"
)

func Test_should_return_error_with_bad_request_response_when_body_is_a_plain_text(t *testing.T) {
	type Type struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	var data Type

	err := lHttp.JsonBody(httpTest.NewRequest(lambler.Json{
		"body": "plain text",
	}), &data)

	response := err.ToResponse().ToLambdaResponse()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.True(t, gfun.EqualsJsonUnmarshal(lambler.Json{
		"message": "bad request",
	}, []byte(response.Body)))
}
