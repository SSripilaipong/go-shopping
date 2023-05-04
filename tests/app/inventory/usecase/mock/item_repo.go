package mock

import inventoryDomain "go-shopping/app/core/inventory/domain"

type ItemRepo struct {
	Create_factory     func() *inventoryDomain.Item
	Create_finalResult *inventoryDomain.Item
	Create_calls       int
}

func (r *ItemRepo) Create(factory func() *inventoryDomain.Item) error {
	calls := r.Create_calls
	if calls == 0 {
		calls = 1
	}

	for i := 0; i < calls; i++ {
		r.Create_finalResult = factory()
	}

	return nil
}
