package lambda

import (
	"go-shopping/app/core/inventory/inventoryController"
)

type dependency struct {
	Inventory inventoryController.Dependency
}

func newDependencyFromConfig(config *Config) dependency {
	return dependency{
		Inventory: inventoryController.NewRuntimeDependency(config.Inventory),
	}
}
