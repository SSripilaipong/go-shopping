package usecase

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

func CreateNewItem() CreateNewItemFunc {
	return func(CreateNewItemRequest) (CreateNewItemResponse, CreateNewItemError) {
		return CreateNewItemResponse{}, nil
	}
}
