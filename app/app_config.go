package app

import (
	inventoryController "go-shopping/app/core/inventory/controller"
	appLambda "go-shopping/app/lambda"
	"go-shopping/lambler"
)

type Config struct {
	Inventory *inventoryController.Config
}

func newConfigFromEnvironment() *Config {
	return &Config{
		Inventory: newInventoryConfigFromEnvironment(),
	}
}

func newDependency(config *Config) appLambda.Dependency {
	return appLambda.Dependency{
		Inventory: inventoryController.NewDependency(lambler.GetDynamodbClient(), config.Inventory),
	}
}
