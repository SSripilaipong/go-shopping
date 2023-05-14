package controller

import (
	"go-shopping/app/core/inventory/inventoryController"
	"go-shopping/app/core/inventory/usecase"
	"go-shopping/app/lambda"
	"go-shopping/lambler"
	"go-shopping/tests"
	"go-shopping/tests/app/inventory/controller/mock"
)

func NewHandlerWithMocks() lambler.Handler {
	return NewHandlerWithInventoryDep(inventoryController.NewDependency(
		(&mock.CreateNewItemUsecase{
			WillReturn: usecase.CreateNewItemResponse{Id: tests.RandomString(16)},
		}).New(),
	))
}

func NewHandlerWithInventoryDep(inventoryDep inventoryController.Dependency) lambler.Handler {
	return lambda.New(inventoryDep)
}

func NewHandlerWithCreateNewItemUsecase(createNewItemUsecase usecase.CreateNewItemFunc) lambler.Handler {
	return NewHandlerWithInventoryDep(inventoryController.NewDependency(createNewItemUsecase))
}
