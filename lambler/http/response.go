package http

import (
	"encoding/json"
)

type Headers = map[string]string

type Response struct {
	StatusCode      int      `json:"statusCode"`
	Headers         Headers  `json:"headers"`
	Body            string   `json:"body"`
	Cookies         []string `json:"cookies"`
	IsBase64Encoded bool     `json:"isBase64Encoded"`
}

func (r *Response) MarshalJSON() ([]byte, error) {
	type Struct Response

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
