package json_body

import (
	"github.com/stretchr/testify/assert"
	gfun "go-shopping/go/func"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
	lTesting "go-shopping/lambler/testing"
	httpTest "go-shopping/tests/lambler/http"
	"net/http"
	"testing"
)

func Test_should_return_error_with_bad_request_response_when_body_is_a_plain_text(t *testing.T) {
	var data struct{}

	err := lHttp.JsonBody(httpTest.NewRequest(lTesting.HttpPostEventWithBody("", "plain text")), &data)

	response := err.ToResponse().ToLambdaResponse()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.True(t, gfun.EqualsJsonUnmarshal(lambler.Json{
		"message": "bad request",
	}, []byte(response.Body)))
}

func Test_should_return_error_with_bad_request_response_when_body_is_absent(t *testing.T) {
	var data struct{}

	err := lHttp.JsonBody(httpTest.NewRequest(lTesting.HttpPostEvent("")), &data)

	response := err.ToResponse().ToLambdaResponse()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.True(t, gfun.EqualsJsonUnmarshal(lambler.Json{
		"message": "bad request",
	}, []byte(response.Body)))
}
