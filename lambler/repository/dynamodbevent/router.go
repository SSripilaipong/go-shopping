package dynamodbevent

type Router interface {
	HandlerFor(record *Record) Handler
	Any(handler Handler)
}

type router struct {
	routes []route
}

type route struct {
	handler Handler
}

func NewRouter() Router {
	return &router{}
}

func (r *router) HandlerFor(*Record) Handler {
	for _, t := range r.routes {
		return t.handler
	}
	return nil
}

func (r *router) Any(handler Handler) {
	r.routes = append(r.routes, route{handler: handler})
}
