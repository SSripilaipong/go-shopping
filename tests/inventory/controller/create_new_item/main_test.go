package create_new_item

import (
	"github.com/stretchr/testify/assert"
	"go-shopping/app/lambda"
	lTesting "go-shopping/lambler/testing"
	"net/http"
	"testing"
)

func Test_should_return_status_created(t *testing.T) {
	raw, _ := lambda.New()(lTesting.NewContext(), lTesting.NewHttpPostEvent("/inventory/v1/items"))
	assert.True(t, lTesting.HttpStatusEquals(raw, http.StatusCreated))
}
