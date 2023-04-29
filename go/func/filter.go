package gfun

import "reflect"

func MapFirstNotZero[T, R any](array []T, mapFunc func(T) R) (r R, ok bool) {
	for _, e := range array {
		r = mapFunc(e)
		if !reflect.ValueOf(r).IsZero() {
			return r, true
		}
	}
	return r, false
}
