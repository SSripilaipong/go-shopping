package http

import (
	"encoding/json"
	"fmt"
)

type Headers = map[string]string

type LambdaResponseBuilder interface {
	ToLambdaResponse() *LambdaResponse
}

type LambdaResponse struct {
	StatusCode      int      `json:"statusCode"`
	Headers         Headers  `json:"headers"`
	Body            string   `json:"body"`
	Cookies         []string `json:"cookies"`
	IsBase64Encoded bool     `json:"isBase64Encoded"`
}

func (r *LambdaResponse) ToLambdaResponse() *LambdaResponse {
	return r
}

func (r *LambdaResponse) MarshalJSON() ([]byte, error) {
	type Struct LambdaResponse

	headers := r.Headers
	if headers == nil {
		headers = make(Headers)
	}

	cookies := r.Cookies
	if cookies == nil {
		cookies = make([]string, 0)
	}

	return json.Marshal(&struct {
		Headers Headers  `json:"headers"`
		Cookies []string `json:"cookies"`
		*Struct
	}{
		Headers: headers,
		Cookies: cookies,
		Struct:  (*Struct)(r),
	})
}

type JsonResponse struct {
	StatusCode int
	Body       any
	Headers    map[string]string
	Cookies    []string
}

func (r JsonResponse) ToLambdaResponse() *LambdaResponse {
	statusCode := r.StatusCode
	if statusCode == 0 {
		statusCode = 200
	}
	body, err := json.Marshal(r.Body)
	if err != nil {
		panic(fmt.Errorf("JsonResponse.ToLambdaResponse() failed to parse body: %w", err))
	}
	return &LambdaResponse{
		StatusCode: statusCode,
		Body:       string(body),
	}
}
