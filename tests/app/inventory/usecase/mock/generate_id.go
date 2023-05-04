package mock

import inventoryUsecase "go-shopping/app/core/inventory/usecase"

type GenerateId struct {
	WillReturn        string
	WillReturns       []string
	willReturns_index int
}

func (g *GenerateId) New() inventoryUsecase.GenerateId {
	return func() string {
		if g.WillReturns != nil {
			index := g.willReturns_index
			g.willReturns_index++
			return g.WillReturns[index]
		}
		return g.WillReturn
	}
}
