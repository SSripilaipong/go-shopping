package inventoryHttp

import "go-shopping/app/core/inventory/usecase"

type Dependency struct {
	CreateNewItemUsecase usecase.CreateNewItemFunc
}
