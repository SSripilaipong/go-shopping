package http

type Router interface {
	HandlerFor(request *Request) Handler
	Post(path string, handler Handler)
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

func (r *router) HandlerFor(*Request) Handler {
	for _, t := range r.routes {
		return func(request *Request) LambdaResponseBuilder {
			return t.handler(request)
		}
	}
	return nil
}

func (r *router) Post(_ string, handler Handler) {
	r.routes = append(r.routes, route{handler: handler})
}
