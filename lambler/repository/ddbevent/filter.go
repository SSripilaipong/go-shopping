package ddbevent

import (
	"context"
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

func newRecord(context.Context, any) (*Record, bool) {
	return &Record{}, true
}

func buildHandler(handler Handler) func(request *Record) (any, error) {
	return func(request *Record) (any, error) {
		handler(request)
		return nil, nil
	}
}
