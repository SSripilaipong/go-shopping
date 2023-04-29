package http

import (
	inventoryUsecase "go-shopping/app/core/inventory/usecase"
)

type Dependency struct {
	CreateNewItemUsecase inventoryUsecase.CreateNewItemFunc
}
