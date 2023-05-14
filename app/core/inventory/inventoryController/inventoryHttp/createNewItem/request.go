package createNewItem

import (
	inventoryUsecase "go-shopping/app/core/inventory/usecase"
	gfun "go-shopping/go/gfunc"
	lHttp "go-shopping/lambler/http"
)

type requestBody struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Quantity    *int    `json:"quantity"`
}

func makeUsecaseRequest(request *lHttp.Request) (*inventoryUsecase.CreateNewItemRequest, lHttp.Error) {
	var body requestBody
	if err := lHttp.JsonBody(request, &body); err != nil {
		return nil, err
	}

	if !gfun.AllStructFieldsSet(body) {
		return nil, lHttp.UnprocessableEntityResponse()
	}

	return &inventoryUsecase.CreateNewItemRequest{
		Name:        *body.Name,
		Description: *body.Description,
		Quantity:    *body.Quantity,
	}, nil
}
