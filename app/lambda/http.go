package lambda

import (
	inventoryHttp "go-shopping/app/core/inventory/http"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
)

func newHttpFilter() lambler.Filter {
	filter := lHttp.NewFilter()
	filter.Include(inventoryHttp.NewRouter())
	return filter
}
