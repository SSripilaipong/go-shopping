package lambda

import (
	inventoryHttp "go-shopping/app/core/inventory/controller/http"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
)

func newHttpFilter(dep Dependency) lambler.Filter {
	filter := lHttp.NewFilter()
	filter.Include(inventoryHttp.NewRouter(dep.Inventory))
	return filter
}
