package router

import (
	"context"
	"net/http"
	"strings"
)
var router *Router

func New() (RouterInterface, error) {
	router = &Router{
		routes: []RouteInterface{},
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
	routes []RouteInterface
}

func (r *Router) RegisterRoutes(routes *[]RouteInterface) {
	for _, route := range *routes {
		r.routes = append(r.routes, route)
	}
}

func matchRoute(pattern, path string) (bool, map[string]string) {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false, nil
	}

	params := make(map[string]string)
	for i, patternPart := range patternParts {
		if strings.HasPrefix(patternPart, ":") {
			params[patternPart[1:]] = pathParts[i]
		} else if patternPart == "*" {
			return true, params
		} else if patternPart != pathParts[i] {
			return false, nil
		}
	}
	return true, params
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	for _, route := range r.routes {
		if match, params := matchRoute(route.GetPath(), req.URL.Path); match {
			if req.Method != route.GetMethod() {
				http.Error(w, "Requested http method is not supported", 405)
				return
			}

			ctx := req.Context()
			for key, value := range params {
				ctx = context.WithValue(ctx, key, value)
			}
			req = req.WithContext(ctx)
			handler := route.GetHandler()
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}