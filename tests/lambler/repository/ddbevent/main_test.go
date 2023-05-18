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
	router.Any(func(record *ddbevent.Record) {
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

func Test_should_pass_record_with_new_image(t *testing.T) {
	type image struct {
		Id string `json:"Id"`
		A  int    `json:"A"`
	}
	var handlerNewImage image

	router := ddbevent.NewRouter()
	router.Any(func(record *ddbevent.Record) {
		_ = ddbevent.NewImageData(record, &handlerNewImage)
	})

	filter := ddbevent.NewFilter()
	filter.Include(router)
	app := lambler.New([]lambler.Filter{filter})

	_, _ = app(context.Background(), lTesting.DynamodbEvent(lTesting.DynamodbEventRecordWithNewImage(
		lambler.Json{"Id": lambler.Json{"S": "abc"}},
		lambler.Json{"Id": lambler.Json{"S": "abc"}, "A": lambler.Json{"N": "123"}},
	)))

	assert.Equal(t, image{
		Id: "abc",
		A:  123,
	}, handlerNewImage)
}
