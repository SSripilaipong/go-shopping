package controller

import (
	inventoryHttp "go-shopping/app/core/inventory/http"
	"go-shopping/app/core/inventory/usecase"
	"go-shopping/app/lambda"
	"go-shopping/lambler"
	"go-shopping/tests"
	"go-shopping/tests/app/inventory/controller/mock"
)

func NewHandlerWithMocks() lambler.Handler {
	handler := NewHandlerWithInventoryDep(
		inventoryHttp.Dependency{
			CreateNewItemUsecase: (&mock.CreateNewItemUsecase{
				WillReturn: usecase.CreateNewItemResponse{Id: tests.RandomString(16)},
			}).New(),
		},
	)
	return handler
}

func NewHandlerWithInventoryDep(inventoryDep inventoryHttp.Dependency) lambler.Handler {
	handler := lambda.New(lambda.Dependency{
		Inventory: inventoryDep,
	})
	return handler
}
