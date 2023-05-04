package create_new_item

import (
	"github.com/stretchr/testify/assert"
	inventoryDomain "go-shopping/app/core/inventory/domain"
	inventoryUsecase "go-shopping/app/core/inventory/usecase"
	"go-shopping/tests/app/inventory/usecase/mock"
	"testing"
)

func Test_should_create_item_in_item_repo_with_generated_id(t *testing.T) {
	itemRepo := &mock.ItemRepo{}
	usecase := inventoryUsecase.CreateNewItem(itemRepo, (&mock.GenerateId{
		WillReturn: "INV-ITM-9999",
	}).New())
	usecase(inventoryUsecase.CreateNewItemRequest{
		Name:        "My Product",
		Description: "My brand-new product",
		Quantity:    555,
	})

	assert.Equal(t, &inventoryDomain.Item{
		Id:          "INV-ITM-9999",
		Name:        "My Product",
		Description: "My brand-new product",
		Quantity:    555,
	}, itemRepo.Create_factory())
}
