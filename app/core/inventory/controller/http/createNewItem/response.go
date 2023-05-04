package createNewItem

import (
	"go-shopping/app/core/inventory/usecase"
	lHttp "go-shopping/lambler/http"
	"net/http"
)

type responseBody struct {
	Id string `json:"id"`
}

func makeHttpResponse(resp usecase.CreateNewItemResponse) lHttp.LambdaResponseBuilder {
	return lHttp.JsonResponse{StatusCode: http.StatusCreated, Body: responseBody{
		Id: resp.Id,
	}}
}
