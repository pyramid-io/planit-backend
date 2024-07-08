package router

import (
	"errors"
	"net/http"
	"strings"
)
var router *Router

func New() (RouterInterface, error) {
	router = &Router{
		routes: make(map[string]RouteInterface),
	}

	return router, nil
}

func GET(path string, handler http.HandlerFunc) Route {
	return Route{
		path: path,
		handler: handler,
		httpMethod: http.MethodGet,
	}
}

func POST(path string, handler http.HandlerFunc) Route {
	return Route{
		path: path,
		handler: handler,
		httpMethod: http.MethodPost,
	}
}

type RouterInterface interface {
	RegisterRoutes(routes *[]RouteInterface)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type RouteInterface interface {
	GetPath() string
	GetHandler() http.HandlerFunc
	GetMethod() string
}

type Route struct {
	path string
	handler http.HandlerFunc
	httpMethod string
}

func (route Route) GetPath() string {
	return route.path
}

func (route Route) GetHandler() http.HandlerFunc {
	return route.handler
}

func (route Route) GetMethod() string {
	return route.httpMethod
}

type Router struct {
	routes map[string]RouteInterface
}

func (r *Router) RegisterRoutes(routes *[]RouteInterface) {
	for _, route := range *routes {
		r.routes[route.GetPath()] = route
	}
}

func matchRoute(pattern, path string) bool {
	if pattern == path {
		return true
	}
	if strings.HasSuffix(pattern, "/*") {
		basePattern := strings.TrimSuffix(pattern, "/*")
		if strings.HasPrefix(path, basePattern) {
			return true
		}
	}
	return false
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if matchRoute(route.GetPath(), req.URL.Path) && req.Method == route.GetMethod() {
			handler := route.GetHandler()
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (r *Router) findRouteByPath(path string) (RouteInterface, error) {
	route, exists := r.routes[path]

	if !exists {
		return nil, errors.New("not found") 
	}

	return route, nil
}
