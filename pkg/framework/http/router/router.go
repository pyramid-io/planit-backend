package router

import (
	"errors"
	"fmt"
	"net/http"
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

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request path is: ", req.URL.Path)
	route, err := r.findRouteByPath(req.URL.Path)
	
	if err != nil {
		fmt.Println("Route is not available for this path: ", req.URL.Path)
		http.NotFound(w, req)
		return
	}

	if route.GetMethod() != req.Method {
		http.Error(w, fmt.Sprintf("Http mehtod (%s) is not supported.", req.Method), http.StatusMethodNotAllowed)	
		return
	} else {
		handler := route.GetHandler()
		handler(w, req)
		return
	}
}

func (r *Router) findRouteByPath(path string) (RouteInterface, error) {
	route, exists := r.routes[path]

	if !exists {
		return nil, errors.New("not found") 
	}

	return route, nil
}
