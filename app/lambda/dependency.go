package lambda

import (
	"go-shopping/app/core/inventory/inventoryController"
	"go-shopping/lambler"
)

type dependency struct {
	Inventory inventoryController.Dependency
}

func newDependencyFromConfig(config *Config) dependency {
	return dependency{
		Inventory: inventoryController.NewRuntimeDependency(lambler.GetDynamodbClient(), config.Inventory),
	}
}
