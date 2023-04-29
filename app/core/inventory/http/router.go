package http

import (
	"go-shopping/app/core/inventory/http/controller"
	lHttp "go-shopping/lambler/http"
)

func NewRouter() lHttp.Router {
	router := lHttp.NewRouter()
	router.Post("/inventory/v1/items", controller.CreateNewItem())
	return router
}
