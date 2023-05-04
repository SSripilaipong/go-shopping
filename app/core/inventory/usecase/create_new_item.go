package usecase

import inventoryDomain "go-shopping/app/core/inventory/domain"

type CreateNewItemRequest struct {
	Name        string
	Description string
	Quantity    int
}
type CreateNewItemResponse struct {
	Id string
}
type CreateNewItemError interface{}

type CreateNewItemFunc = func(CreateNewItemRequest) (CreateNewItemResponse, CreateNewItemError)

func CreateNewItem(itemRepo ItemRepo, generateId GenerateId) CreateNewItemFunc {
	return func(request CreateNewItemRequest) (CreateNewItemResponse, CreateNewItemError) {
		id := generateId()
		if err := itemRepo.Create(func() *inventoryDomain.Item {
			return &inventoryDomain.Item{
				Id:          id,
				Name:        request.Name,
				Description: request.Description,
				Quantity:    request.Quantity,
			}
		}); err != nil {
			panic(err)
		}

		return CreateNewItemResponse{
			Id: id,
		}, nil
	}
}
