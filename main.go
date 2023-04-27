package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	lHandler "go-shopping/lambler/handler"
)

func main() {
	//lambda.Start(lambda.Handler)
	lambda.Start(lHandler.Echo)
}
