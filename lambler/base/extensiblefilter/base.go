package extensiblefilter

import (
	"context"
	gfun "go-shopping/go/gfunc"
)

type Base[D any, H any, R router[D, H]] struct {
	routers      []router[D, H]
	transform    func(context.Context, any) (D, bool)
	buildHandler func(H) func(D) (any, error)
}

func New[D any, H any, R router[D, H]](
	transform func(context.Context, any) (D, bool),
	buildHandler func(H) func(D) (any, error),
) *Base[D, H, R] {
	return &Base[D, H, R]{
		transform:    transform,
		buildHandler: buildHandler,
	}
}

func (f *Base[D, H, R]) Filter(ctx context.Context, event any) func(context.Context, any) (any, error) {
	data, success := f.transform(ctx, event)
	if !success {
		return nil
	}

	handler, found := gfun.MapFirstNotZero(f.routers, func(t router[D, H]) H { return t.HandlerFor(data) })
	if !found {
		return nil
	}

	return func(ctx context.Context, a any) (any, error) {
		return f.buildHandler(handler)(data)
	}
}

func (f *Base[D, H, R]) Include(router R) {
	f.routers = append(f.routers, router)
}
