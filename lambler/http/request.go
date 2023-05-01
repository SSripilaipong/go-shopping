package http

import (
	"context"
	"encoding/json"
	"go-shopping/lambler"
)

type Request struct {
	event any
}

func JsonBody[T any](request *Request, target *T) Error {
	event, isJson := request.event.(lambler.Json)
	if !isJson {
		panic("invalid lambda event")
	}
	bodyAny := event["body"]
	bodyBytes, isString := bodyAny.(string)
	if !isString {
		panic("bad request")
	}
	if err := json.Unmarshal([]byte(bodyBytes), target); err != nil {
		return UnprocessableEntityResponse()
	}
	return nil
}

func NewRequest(_ context.Context, event any) (*Request, error) {
	return &Request{event: event}, nil
}
