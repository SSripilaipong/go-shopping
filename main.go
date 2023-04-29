package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	appLambda "go-shopping/app/lambda"
)

func main() {
	lambda.Start(appLambda.New())
	//lambda.Start(lHandler.Echo)
}
