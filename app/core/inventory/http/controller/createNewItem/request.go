package createNewItem

import (
	"fmt"
	inventoryUsecase "go-shopping/app/core/inventory/usecase"
	lHttp "go-shopping/lambler/http"
)

type jsonBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func makeUsecaseRequest(request *lHttp.Request) *inventoryUsecase.CreateNewItemRequest {
	var body jsonBody
	err := lHttp.JsonBody(request, &body)
	if err != nil {
		panic(fmt.Errorf("unhandled error: %w", err))
	}

	return &inventoryUsecase.CreateNewItemRequest{
		Name:        body.Name,
		Description: body.Description,
		Quantity:    body.Quantity,
	}
}
