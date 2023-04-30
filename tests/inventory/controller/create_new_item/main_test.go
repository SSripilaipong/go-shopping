package create_new_item

import (
	"github.com/stretchr/testify/assert"
	inventoryHttp "go-shopping/app/core/inventory/http"
	inventoryUsecase "go-shopping/app/core/inventory/usecase"
	"go-shopping/lambler"
	lTesting "go-shopping/lambler/testing"
	controllerTest "go-shopping/tests/inventory/controller"
	"go-shopping/tests/inventory/controller/mock"
	"net/http"
	"testing"
)

func Test_should_return_status_created(t *testing.T) {
	handler := controllerTest.NewHandlerWithMocks()

	raw, _ := doValidRequest(handler)

	assert.Equal(t, lTesting.NewTestHttpResponse(raw).StatusCode(), http.StatusCreated)
}

func Test_should_execute_create_new_item_usecase_with_payload(t *testing.T) {
	usecase := &mock.CreateNewItemUsecase{}
	handler := controllerTest.NewHandlerWithInventoryDep(inventoryHttp.Dependency{
		CreateNewItemUsecase: usecase.New(),
	})

	_, _ = doValidRequestWithPayload(handler, validPayload("Product A", "Good One", 999))

	assert.Equal(t, inventoryUsecase.CreateNewItemRequest{
		Name:        "Product A",
		Description: "Good One",
		Quantity:    999,
	}, usecase.Request)
}

func Test_should_respond_with_item_id(t *testing.T) {
	handler := controllerTest.NewHandlerWithInventoryDep(inventoryHttp.Dependency{
		CreateNewItemUsecase: (&mock.CreateNewItemUsecase{
			WillReturn: inventoryUsecase.CreateNewItemResponse{Id: "INV-ITM-1234"},
		}).New(),
	})

	raw, _ := doValidRequest(handler)

	assert.Equal(t, lTesting.NewTestHttpResponse(raw).Json(), lambler.Json{
		"id": "INV-ITM-1234",
	})
}
