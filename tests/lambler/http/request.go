package http

import (
	"context"
	"fmt"
	lHttp "go-shopping/lambler/http"
)

func NewRequest(event any) *lHttp.Request {
	req, success := lHttp.NewRequest(context.Background(), event)
	if !success {
		panic(fmt.Errorf("cannot transform event to request"))
	}
	return req
}
