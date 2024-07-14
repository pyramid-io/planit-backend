package routes

import (
	"fmt"
	"net/http"

	_ "github.com/pyramid.io/planit-backend/docs"
	"github.com/pyramid.io/planit-backend/internal_module/handlers/dummy_handler"
	json_handler "github.com/pyramid.io/planit-backend/internal_module/handlers/json_handler"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetRoutes() *[]router.RouteInterface {
	return &routes
}

var routes = []router.RouteInterface{
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the Home Page")
	}),
	router.GET("/ping", dummy_handler.Ping),
	router.GET("/pong", dummy_handler.Pong),
	router.GET("/json", json_handler.JsonResponseHandler),
	router.GET("/test-logger", dummy_handler.TestLogger),
	router.GET("/time", dummy_handler.TimeHandler),
	router.GET("/with-params/:i/:j", dummy_handler.WithParams),
	router.GET("/swagger/*", httpSwagger.WrapHandler),
	router.GET("/session-create", dummy_handler.SessionCreateHandler),
	router.GET("/session-get", dummy_handler.SessionGetHandler),
	router.GET("/route-with-middleware", dummy_handler.MiddlewareTest).
		Middlewares(
			LoggingMiddlewareInternal,
		),
}
