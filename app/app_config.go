package app

import (
	appLambda "go-shopping/app/lambda"
	"go-shopping/lambler"
)

type Config struct {
	Inventory *InventoryConfig
}

func newConfigFromEnvironment() *Config {
	return &Config{
		Inventory: newInventoryConfigFromEnvironment(),
	}
}

func newDependency(config *Config) appLambda.Dependency {
	return appLambda.Dependency{
		Inventory: newInventoryDependency(lambler.GetDynamodbClient(), config.Inventory),
	}
}
