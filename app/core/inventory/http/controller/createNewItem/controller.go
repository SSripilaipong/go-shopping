package createNewItem

import (
	"go-shopping/app/core/inventory/usecase"
	lHttp "go-shopping/lambler/http"
	"net/http"
)

func New(execute usecase.CreateNewItemFunc) lHttp.Handler {
	return func(request *lHttp.Request) *lHttp.Response {
		usecaseRequest := *makeUsecaseRequest(request)

		execute(usecaseRequest)

		return &lHttp.Response{StatusCode: http.StatusCreated}
	}
}
