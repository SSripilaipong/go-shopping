package controller

import (
	lHttp "go-shopping/lambler/http"
	"net/http"
)

func CreateNewItem() lHttp.Handler {
	return func(request *lHttp.Request) *lHttp.Response {
		return &lHttp.Response{StatusCode: http.StatusCreated}
	}
}
