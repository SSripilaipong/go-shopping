package app

import (
	"go-shopping/app/core/inventory/inventoryController"
	appLambda "go-shopping/app/lambda"
	"os"
)

func newConfigFromEnvironment() *appLambda.Config {
	return &appLambda.Config{
		Inventory: newInventoryConfigFromEnvironment(),
	}
}

func newInventoryConfigFromEnvironment() *inventoryController.Config {
	return &inventoryController.Config{
		ItemTableName: os.Getenv("INVENTORY_ITEM_TABLE_NAME"),
	}
}
