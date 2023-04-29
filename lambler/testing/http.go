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
	event["body"] = bodyJson
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

func NewHttpResponseFromJson(raw lambler.Json) (resp lHttp.Response, err error) {
	if err = gfun.UnmarshalFromMap(raw, &resp); err != nil {
		return lHttp.Response{}, fmt.Errorf("NewHttpResponseFromJson failed to unmarshal data: %w", err)
	}
	return resp, nil
}

func HttpStatusEquals(raw any, status int) bool {
	data, ok := raw.(lambler.Json)
	if !ok {
		return false
	}
	resp, err := NewHttpResponseFromJson(data)
	if err != nil {
		return false
	}
	return resp.StatusCode == status
}
