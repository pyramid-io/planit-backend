package router

import (
	"net/http"
)

type RouterInterface interface {
	RegisterRoutes(routes *[]Route)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type Router struct {
	routes map[string]http.HandlerFunc
}

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

func New() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) RegisterRoutes(routes *[]Route) {
	for _, route := range *routes {
		r.routes[route.Path] = route.Handler
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, exists := r.routes[req.URL.Path]; exists {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}
