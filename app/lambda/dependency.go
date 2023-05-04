package lambda

import (
	inventoryHttp "go-shopping/app/core/inventory/controller/http"
)

type Dependency struct {
	Inventory inventoryHttp.Dependency
}
