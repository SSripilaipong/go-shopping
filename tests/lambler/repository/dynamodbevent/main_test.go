package dynamodbevent

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-shopping/lambler"
	"go-shopping/lambler/repository/dynamodbevent"
	lTesting "go-shopping/lambler/testing"
	"testing"
)

func Test_should_call_handler(t *testing.T) {
	var handlerIsCalled bool

	router := dynamodbevent.NewRouter()
	router.Any(func(request *dynamodbevent.Record) {
		handlerIsCalled = true
	})

	filter := dynamodbevent.NewFilter()
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
