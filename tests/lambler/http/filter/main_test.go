package filter

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
	"net/http"
	"testing"
)

func Test_should_not_handle_event_with_unknown_structure(t *testing.T) {
	var handlerIsCalled bool
	router := lHttp.NewRouter()
	router.Post("", func(request *lHttp.Request) lHttp.LambdaResponseBuilder {
		handlerIsCalled = true
		return lHttp.JsonResponse{StatusCode: http.StatusCreated}
	})

	filter := lHttp.NewFilter()
	filter.Include(router)
	app := lambler.New([]lambler.Filter{filter})

	_, _ = app(context.Background(), lambler.Json{})

	assert.False(t, handlerIsCalled)
}
