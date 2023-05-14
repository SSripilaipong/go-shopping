package http

import (
	"context"
	"encoding/json"
	"fmt"
	gfun "go-shopping/go/gfunc"
	"go-shopping/lambler"
)

type LambdaFunctionURLRequest struct {
	Version               *string                          `json:"version" validate:"required"` // Version is expected to be `"2.0"`
	RawPath               *string                          `json:"rawPath" validate:"required"`
	RawQueryString        *string                          `json:"rawQueryString"`
	Cookies               []string                         `json:"cookies,omitempty"`
	Headers               map[string]string                `json:"headers"`
	QueryStringParameters map[string]string                `json:"queryStringParameters,omitempty"`
	Body                  *string                          `json:"body,omitempty"`
	IsBase64Encoded       *bool                            `json:"isBase64Encoded"`
	RequestContext        *LambdaFunctionURLRequestContext `json:"requestContext" validate:"required"`
}

type LambdaFunctionURLRequestContext struct {
	Http *LambdaFunctionURLRequestContextHTTPDescription `json:"http" validate:"required"`
}

type LambdaFunctionURLRequestContextHTTPDescription struct {
	Method *string `json:"method" validate:"required"`
	Path   *string `json:"path" validate:"required"`
}

type Request struct {
	event *LambdaFunctionURLRequest
}

func JsonBody[T any](request *Request, target *T) Error {
	body := request.event.Body
	if body == nil {
		return BadRequestResponse()
	}
	if err := json.Unmarshal([]byte(*body), target); err != nil {
		return BadRequestResponse()
	}
	return nil
}

func NewRequest(_ context.Context, event any) (*Request, error) {
	eventJson, isJson := event.(lambler.Json)
	if !isJson {
		return nil, fmt.Errorf("invalid event type")
	}

	unmarshalledEvent, err := validateEvent(eventJson)
	if err != nil {
		return nil, err
	}
	return &Request{event: unmarshalledEvent}, nil
}

func validateEvent(eventJson lambler.Json) (*LambdaFunctionURLRequest, error) {
	var unmarshalledEvent LambdaFunctionURLRequest
	if err := gfun.UnmarshalFromMap(eventJson, &unmarshalledEvent); err != nil {
		return nil, err
	}
	if err := gfun.ValidateStruct(unmarshalledEvent); err != nil {
		return nil, err
	}
	return &unmarshalledEvent, nil
}
