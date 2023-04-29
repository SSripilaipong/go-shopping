package mock

import inventoryUsecase "go-shopping/app/core/inventory/usecase"

type CreateNewItemUsecase struct {
	Request inventoryUsecase.CreateNewItemRequest
}

func (u *CreateNewItemUsecase) New() inventoryUsecase.CreateNewItemFunc {
	return func(request inventoryUsecase.CreateNewItemRequest) (inventoryUsecase.CreateNewItemResponse, inventoryUsecase.CreateNewItemError) {
		u.Request = request
		return inventoryUsecase.CreateNewItemResponse{}, nil
	}
}
