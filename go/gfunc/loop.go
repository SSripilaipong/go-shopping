package gfun

func ForEach[T any](array []T, f func(T)) {
	for _, x := range array {
		f(x)
	}
}
