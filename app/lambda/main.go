package lambda

import (
	"go-shopping/lambler"
)

func New(dep Dependency) lambler.Handler {
	return lambler.New([]lambler.Filter{
		newHttpFilter(dep),
		lambler.NewPrintFilter(),
	})
}
