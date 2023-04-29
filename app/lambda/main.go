package lambda

import (
	"go-shopping/lambler"
)

func New() lambler.Handler {
	return lambler.New([]lambler.Filter{
		newHttpFilter(),
	})
}
