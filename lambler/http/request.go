package http

import (
	"context"
)

type Request struct {
}

func NewRequest(_ context.Context, _ any) (*Request, error) {
	return &Request{}, nil
}
