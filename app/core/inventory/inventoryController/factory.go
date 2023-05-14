package inventoryController

import (
	"go-shopping/app/core/inventory/inventoryController/inventoryHttp"
	lHttp "go-shopping/lambler/http"
)

type Factory struct {
	dep Dependency
}

func NewFactory(dep Dependency) *Factory {
	return &Factory{dep: dep}
}

func (f *Factory) NewHttpRouter() lHttp.Router {
	return inventoryHttp.NewRouter(f.dep.http)
}
