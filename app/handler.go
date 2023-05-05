package app

import (
	appLambda "go-shopping/app/lambda"
	"go-shopping/lambler"
)

func NewAppHandler() lambler.Handler {
	return appLambda.New(newDependency(newConfigFromEnvironment()))
}
