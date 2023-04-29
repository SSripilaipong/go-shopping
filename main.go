package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	inventoryHttp "go-shopping/app/core/inventory/http"
	appLambda "go-shopping/app/lambda"
)

func main() {
	lambda.Start(appLambda.New(appLambda.Dependency{
		Inventory: inventoryHttp.Dependency{},
	}))
	//lambda.Start(lambler.EchoHandler)
}
