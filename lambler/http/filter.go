package http

import (
	gfun "go-shopping/go/gfunc"
	"go-shopping/lambler"
	"go-shopping/lambler/base/extensiblefilter"
)

type Filter interface {
	lambler.Filter
	Include(router Router)
}

func NewFilter() Filter {
	return extensiblefilter.New[*Request, Handler, Router](NewRequest, buildHandler)
}

func buildHandler(handler Handler) func(request *Request) (any, error) {
	return func(request *Request) (any, error) {
		return gfun.MarshalToMap(handler(request).ToLambdaResponse())
	}
}
