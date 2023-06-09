package inventoryHttp

import (
	"go-shopping/app/core/inventory/inventoryController/inventoryHttp/createNewItem"
	lHttp "go-shopping/lambler/http"
)

func NewRouter(dep Dependency) lHttp.Router {
	router := lHttp.NewRouter()
	router.Post("/inventory/v1/items", createNewItem.New(dep.CreateNewItemUsecase))
	return router
}
