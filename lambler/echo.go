package lambler

import (
	"context"
)

func EchoHandler(_ context.Context, event any) (any, error) {
	return event, nil
}
