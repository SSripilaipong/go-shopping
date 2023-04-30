package createNewItem

import (
	"go-shopping/app/core/inventory/usecase"
	lHttp "go-shopping/lambler/http"
)

func New(execute usecase.CreateNewItemFunc) lHttp.Handler {
	return func(request *lHttp.Request) lHttp.LambdaResponseBuilder {
		usecaseRequest := *makeUsecaseRequest(request)

		resp, _ := execute(usecaseRequest)

		return makeHttpResponse(resp)
	}
}
