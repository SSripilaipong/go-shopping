package testing

import (
	"fmt"
	gfun "go-shopping/go/func"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
)

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

func HttpStatusIs(raw lambler.Json, status int) bool {
	resp, err := NewHttpResponseFromJson(raw)
	if err != nil {
		return false
	}
	return resp.StatusCode == status
}
