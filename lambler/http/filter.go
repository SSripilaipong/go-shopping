package http

import (
	"context"
	gfun "go-shopping/go/func"
	"go-shopping/lambler"
)

type Filter interface {
	lambler.Filter
	Include(router Router)
}

type filter struct {
	routers []Router
}

func NewFilter() Filter {
	return &filter{}
}

func (f *filter) Filter(ctx context.Context, event any) lambler.Handler {
	request, err := NewRequest(ctx, event)
	if err != nil {
		return nil
	}

	handler, found := gfun.MapFirstNotZero(f.routers, func(t Router) Handler { return t.HandlerFor(request) })
	if !found {
		return nil
	}

	return func(context.Context, any) (any, error) {
		return gfun.MarshalToMap(handler(request).ToLambdaResponse())
	}
}

func (f *filter) Include(router Router) {
	f.routers = append(f.routers, router)
}
