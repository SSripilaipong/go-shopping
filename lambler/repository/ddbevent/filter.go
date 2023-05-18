package ddbevent

import (
	"context"
	"fmt"
	gfun "go-shopping/go/gfunc"
	"go-shopping/lambler"
	"go-shopping/lambler/base/extensiblefilter"
)

type Filter interface {
	lambler.Filter
	Include(router Router)
}

func NewFilter() Filter {
	return extensiblefilter.New[*Record, Handler, Router](newRecord, buildHandler)
}

func newRecord(_ context.Context, event any) (*Record, bool) {
	mapEvent, ok := event.(map[string]any)
	if !ok {
		panic(fmt.Errorf("unhandled error"))
	}

	var parsedEvent dynamoDBEvent
	err := gfun.UnmarshalFromMap(mapEvent, &parsedEvent)
	if err != nil {
		panic(fmt.Errorf("unhandled error"))
	}

	return &Record{newImage: parsedEvent.Records[0].Change.NewImage}, true
}

func buildHandler(handler Handler) func(request *Record) (any, error) {
	return func(request *Record) (any, error) {
		handler(request)
		return nil, nil
	}
}
