package lambler

import (
	"context"
)

type Filter interface {
	Filter(ctx context.Context, event any) Handler
}
