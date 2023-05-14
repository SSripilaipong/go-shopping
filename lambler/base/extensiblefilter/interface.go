package extensiblefilter

import (
	"context"
)

type Interface[D any, H any, R router[D, H]] interface {
	Filter(context.Context, any) func(context.Context, any) (any, error)
	Include(router R)
}

type router[D any, H any] interface {
	HandlerFor(D) H
}
