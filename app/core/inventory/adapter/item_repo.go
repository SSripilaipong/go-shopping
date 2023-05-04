package adapter

import (
	inventoryDomain "go-shopping/app/core/inventory/domain"
	"go-shopping/app/core/inventory/usecase"
)

func NewItemRepo() usecase.ItemRepo {
	return &itemRepo{}
}

type itemRepo struct {
}

func (i *itemRepo) Create(factory func() *inventoryDomain.Item) error {
	factory()
	return nil
}
