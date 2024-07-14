package routes

import (
	"fmt"
	"net/http"
)

func LoggingMiddlewareInternal(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging middleware Internal: URL =", r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}