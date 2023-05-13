package lambler

import (
	"context"
	"fmt"
)

type Filter interface {
	Filter(ctx context.Context, event any) Handler
}

type printFilter struct {
}

func (printFilter) Filter(context.Context, any) Handler {
	return func(_ context.Context, event any) (any, error) {
		fmt.Printf("%#v\n", event)
		return nil, nil
	}
}

func NewPrintFilter() Filter {
	return printFilter{}
}
