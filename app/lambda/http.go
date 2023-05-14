package lambda

import (
	"go-shopping/go/gfunc"
	"go-shopping/lambler"
	lHttp "go-shopping/lambler/http"
)

func newHttpFilter(routers ...lHttp.Router) lambler.Filter {
	filter := lHttp.NewFilter()
	gfun.ForEach(routers, func(router lHttp.Router) {
		filter.Include(router)
	})
	return filter
}
