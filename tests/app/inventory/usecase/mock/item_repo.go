package mock

import inventoryDomain "go-shopping/app/core/inventory/domain"

type ItemRepo struct {
	Create_factory func() *inventoryDomain.Item
}

func (r *ItemRepo) Create(factory func() *inventoryDomain.Item) error {
	r.Create_factory = factory
	return nil
}
