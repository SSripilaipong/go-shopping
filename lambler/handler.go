package lambler

import "context"

type Handler = func(ctx context.Context, event any) (any, error)

func New(filters []Filter) Handler {
	return func(ctx context.Context, event any) (any, error) {
		for _, filter := range filters {
			execute := filter.Filter(ctx, event)
			if execute != nil {
				return execute(ctx, event)
			}
		}
		return nil, nil
	}
}
