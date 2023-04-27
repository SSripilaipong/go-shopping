package create_new_item

import (
	"github.com/stretchr/testify/assert"
	"go-shopping/app/lambda"
	"go-shopping/lambler"
	lTesting "go-shopping/lambler/testing"
	"net/http"
	"testing"
)

func Test_should_return_status_created(t *testing.T) {
	raw, _ := lambda.Handler(lTesting.NewContext(), lTesting.NewHttpPostEvent("/inventory/v1/item"))

	assert.True(t, lTesting.HttpStatusIs(raw.(lambler.Json), http.StatusCreated))
}
