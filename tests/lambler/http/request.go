package http

import (
	"context"
	lHttp "go-shopping/lambler/http"
)

func NewRequest(event any) *lHttp.Request {
	req, err := lHttp.NewRequest(context.Background(), event)
	if err != nil {
		panic(err)
	}
	return req
}
