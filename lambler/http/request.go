package http

import (
	"context"
	"encoding/json"
	"fmt"
	"go-shopping/lambler"
)

type Request struct {
	event any
}

func JsonBody[T any](request *Request, target *T) error {
	event, isJson := request.event.(lambler.Json)
	if !isJson {
		panic("invalid lambda event")
	}
	bodyAny := event["body"]
	bodyString, isString := bodyAny.([]byte)
	if !isString {
		return fmt.Errorf("unprocessable entity")
	}
	if err := json.Unmarshal(bodyString, target); err != nil {
		return fmt.Errorf("unprocessable entity")
	}
	return nil
}

func NewRequest(_ context.Context, event any) (*Request, error) {
	return &Request{event: event}, nil
}
