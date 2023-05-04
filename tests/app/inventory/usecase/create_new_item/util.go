package create_new_item

import (
	inventoryUsecase "go-shopping/app/core/inventory/usecase"
	"go-shopping/tests"
	"go-shopping/tests/app/inventory/usecase/mock"
	"math/rand"
)

func newUsecaseWithGenerateId(generateId inventoryUsecase.GenerateId) inventoryUsecase.CreateNewItemFunc {
	return newUsecaseWithItemRepoAndGenerateId(&mock.ItemRepo{}, generateId)
}

func newUsecaseWithItemRepoAndGenerateId(itemRepo inventoryUsecase.ItemRepo, generateId inventoryUsecase.GenerateId) inventoryUsecase.CreateNewItemFunc {
	return inventoryUsecase.CreateNewItem(itemRepo, generateId)
}

func doExecuteWithRequest(usecase inventoryUsecase.CreateNewItemFunc, request inventoryUsecase.CreateNewItemRequest) (inventoryUsecase.CreateNewItemResponse, inventoryUsecase.CreateNewItemError) {
	return usecase(request)
}

func doValidExecute(usecase inventoryUsecase.CreateNewItemFunc) (inventoryUsecase.CreateNewItemResponse, inventoryUsecase.CreateNewItemError) {
	return doExecuteWithRequest(usecase, inventoryUsecase.CreateNewItemRequest{
		Name:        tests.RandomString(16),
		Description: tests.RandomString(64),
		Quantity:    rand.Intn(100),
	})
}
