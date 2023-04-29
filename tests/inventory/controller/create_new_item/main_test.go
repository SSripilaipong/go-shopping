package create_new_item

import (
	"github.com/stretchr/testify/assert"
	inventoryHttp "go-shopping/app/core/inventory/http"
	inventoryUsecase "go-shopping/app/core/inventory/usecase"
	lTesting "go-shopping/lambler/testing"
	controllerTest "go-shopping/tests/inventory/controller"
	"go-shopping/tests/inventory/controller/mock"
	"net/http"
	"testing"
)

func Test_should_return_status_created(t *testing.T) {
	handler := controllerTest.NewHandlerWithMocks()

	raw, _ := doValidRequest(handler)

	assert.True(t, lTesting.HttpStatusEquals(raw, http.StatusCreated))
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
