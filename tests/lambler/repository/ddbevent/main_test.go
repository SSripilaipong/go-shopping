package ddbevent

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-shopping/lambler"
	"go-shopping/lambler/repository/ddbevent"
	lTesting "go-shopping/lambler/testing"
	"testing"
)

func Test_should_call_handler(t *testing.T) {
	var handlerIsCalled bool

	router := ddbevent.NewRouter()
	router.Any(func(request *ddbevent.Record) {
		handlerIsCalled = true
	})

	filter := ddbevent.NewFilter()
	filter.Include(router)
	app := lambler.New([]lambler.Filter{filter})

	_, _ = app(context.Background(), lTesting.DynamodbEvent(
		lTesting.DynamodbEventRecordWithNewImage(
			lambler.Json{"Id": lambler.Json{"S": "abc"}},
			lambler.Json{"Id": lambler.Json{"S": "abc"}},
		),
	))

	assert.True(t, handlerIsCalled)
}
