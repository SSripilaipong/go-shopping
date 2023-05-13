package app

import (
	"go-shopping/app/core/inventory/controller"
	"os"
)

func newInventoryConfigFromEnvironment() *controller.Config {
	return &controller.Config{
		ItemTableName: os.Getenv("INVENTORY_ITEM_TABLE_NAME"),
	}
}
