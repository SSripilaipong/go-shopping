package handler

import (
	"context"
)

func Echo(_ context.Context, event any) (any, error) {
	return event, nil
}
