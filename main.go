package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	inventoryHttp "go-shopping/app/core/inventory/http"
	"go-shopping/app/core/inventory/usecase"
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
			CreateNewItemUsecase: usecase.CreateNewItem(),
		},
	})
}
