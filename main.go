package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	inventoryAdapter "go-shopping/app/core/inventory/adapter"
	inventoryHttp "go-shopping/app/core/inventory/controller/http"
	inventoryUsecase "go-shopping/app/core/inventory/usecase"
	appLambda "go-shopping/app/lambda"
	"go-shopping/lambler"
)

func main() {
	lambda.Start(NewAppHandler())
	//lambda.Start(lambler.EchoHandler)
}

func NewAppHandler() lambler.Handler {
	return appLambda.New(appLambda.Dependency{
		Inventory: inventoryHttp.Dependency{
			CreateNewItemUsecase: inventoryUsecase.CreateNewItem(
				inventoryAdapter.NewItemRepo(),
				inventoryAdapter.GenerateItemId,
			),
		},
	})
}
