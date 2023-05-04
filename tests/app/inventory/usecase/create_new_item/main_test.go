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
	usecase := newUsecaseWithItemRepoAndGenerateId(itemRepo, (&mock.GenerateId{
		WillReturn: "INV-ITM-9999",
	}).New())

	doExecuteWithRequest(usecase, inventoryUsecase.CreateNewItemRequest{
		Name:        "My Product",
		Description: "My brand-new product",
		Quantity:    555,
	})

	assert.Equal(t, &inventoryDomain.Item{
		Id:          "INV-ITM-9999",
		Name:        "My Product",
		Description: "My brand-new product",
		Quantity:    555,
	}, itemRepo.Create_finalResult)
}

func Test_should_return_response_with_generated_id(t *testing.T) {
	usecase := newUsecaseWithGenerateId((&mock.GenerateId{
		WillReturn: "INV-ITM-9999",
	}).New())

	response, _ := doValidExecute(usecase)

	assert.Equal(t, inventoryUsecase.CreateNewItemResponse{Id: "INV-ITM-9999"}, response)
}

func Test_should_return_response_with_latest_generated_id(t *testing.T) {
	usecase := newUsecaseWithItemRepoAndGenerateId(&mock.ItemRepo{
		Create_calls: 2,
	}, (&mock.GenerateId{
		WillReturns: []string{"INV-ITM-1234", "INV-ITM-5678"},
	}).New())

	response, _ := doValidExecute(usecase)

	assert.Equal(t, inventoryUsecase.CreateNewItemResponse{Id: "INV-ITM-5678"}, response)
}
