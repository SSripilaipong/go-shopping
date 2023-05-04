package mock

import inventoryUsecase "go-shopping/app/core/inventory/usecase"

type GenerateId struct {
	WillReturn string
}

func (g *GenerateId) New() inventoryUsecase.GenerateId {
	return func() string {
		return g.WillReturn
	}
}
