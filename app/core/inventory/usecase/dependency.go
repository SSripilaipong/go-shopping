package usecase

import inventoryDomain "go-shopping/app/core/inventory/domain"

type ItemRepo interface {
	Create(factory func() *inventoryDomain.Item) error
}

type GenerateId func() string
