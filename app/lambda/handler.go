package lambda

import (
	"context"
	gfun "go-shopping/go/func"
	lHttp "go-shopping/lambler/http"
	"net/http"
)

func Handler(_ context.Context, event any) (any, error) {
	resp, err := gfun.MarshalToMap(&lHttp.Response{
		StatusCode: http.StatusCreated,
	})
	return resp, err
}
