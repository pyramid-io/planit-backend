package router

import (
	"net/http"
)

type RouterInterface interface {
	RegisterRoutes(routes *[]Route)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type router struct {
	routes map[string]http.HandlerFunc
}

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

func New() *router {
	return &router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *router) RegisterRoutes(routes *[]Route) {
	for _, route := range *routes {
		r.routes[route.Path] = route.Handler
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, exists := r.routes[req.URL.Path]; exists {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}
