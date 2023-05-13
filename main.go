package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"go-shopping/app"
)

func main() {
	lambda.Start(app.NewAppHandler())
	//lambda.Start(lambler.PrintHandler)
}
