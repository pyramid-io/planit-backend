package middleware

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

var BasicMiddlewares = []Middleware{
	RouteLoggerMiddleware,
}

func RouteLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("incoming route is: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}