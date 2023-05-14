package lambda

import (
	"go-shopping/app/core/inventory/inventoryController"
	"go-shopping/go/gfunc"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
)

type controllerFactory interface {
	NewHttpRouter() lHttp.Router
}

func New(inventoryDep inventoryController.Dependency) lambler.Handler {
	controllers := []controllerFactory{
		inventoryController.NewFactory(inventoryDep),
	}

	httpRouters := gfun.Map(controllers, func(t controllerFactory) lHttp.Router {
		return t.NewHttpRouter()
	})

	return lambler.New([]lambler.Filter{
		newHttpFilter(httpRouters...),
		lambler.NewPrintFilter(),
	})
}

func NewFromConfig(config *Config) lambler.Handler {
	dep := newDependencyFromConfig(config)
	return New(dep.Inventory)
}
