package gfun

func Pointer[T any](x T) *T {
	return &x
}
