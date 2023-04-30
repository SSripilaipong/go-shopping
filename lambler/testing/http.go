package testing

import (
	"encoding/json"
	"fmt"
	gfun "go-shopping/go/func"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
)

func HttpPostEventWithJsonBody(url string, body lambler.Json) lambler.Json {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	event := NewHttpPostEvent(url)
	event["body"] = string(bodyJson)
	return event
}

func NewHttpPostEvent(url string) lambler.Json {
	return lambler.Json{
		"version": "2.0",
		"rawPath": url,
		"requestContext": lambler.Json{
			"http": lambler.Json{
				"path":   url,
				"method": "POST",
			},
		},
	}
}

type TestHttpResponse struct {
	statusCode int
	body       string
}

func (r TestHttpResponse) StatusCode() int {
	return r.statusCode
}

func (r TestHttpResponse) Json() (jsonBody lambler.Json) {
	err := json.Unmarshal([]byte(r.body), &jsonBody)
	if err != nil {
		panic(fmt.Errorf("TestHttpResponse.Json() failed to parse body to json: %w", err))
	}
	return jsonBody
}

func NewTestHttpResponse(raw any) TestHttpResponse {
	jsonObj, ok := raw.(lambler.Json)
	if !ok {
		panic(fmt.Errorf("NewTestHttpResponse failed to unmarshal data: cannot cast data to json"))
	}

	fmt.Printf("%#v\n", jsonObj)
	var resp lHttp.LambdaResponse
	if err := gfun.UnmarshalFromMap(jsonObj, &resp); err != nil {
		panic(fmt.Errorf("NewTestHttpResponse failed to unmarshal data: %w", err))
	}

	return TestHttpResponse{
		statusCode: resp.StatusCode,
		body:       resp.Body,
	}
}
