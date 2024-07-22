package routes

import (
	_ "github.com/pyramid.io/planit-backend/docs"
	"github.com/pyramid.io/planit-backend/internal_module/handlers/dummy_handler"
	json_handler "github.com/pyramid.io/planit-backend/internal_module/handlers/json_handler"
	"github.com/pyramid.io/planit-backend/pkg/framework/application"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetRoutes() *[]router.RouteInterface {
	return &routes
}

var routes = []router.RouteInterface{
	router.GET("/", dummy_handler.Home),
	router.GET("/ping", dummy_handler.Ping),
	router.GET("/pong", dummy_handler.Pong),
	router.GET("/json", json_handler.JsonResponseHandler),
	router.GET("/test-logger", dummy_handler.TestLogger),
	router.GET("/time", dummy_handler.TimeHandler),
	router.GET("/with-params/:i/:j", dummy_handler.WithParams),
	router.GET("/swagger/*", httpSwagger.WrapHandler),
	router.GET("/session-create", dummy_handler.SessionCreateHandler),
	router.GET("/session-get", dummy_handler.SessionGetHandler),
	router.GET("/database-test", dummy_handler.DatabaseTest),
	router.GET("/route-with-middleware", dummy_handler.MiddlewareTest).
		Middlewares(
			LoggingMiddlewareInternal,
		),
	router.GET("/login", dummy_handler.Login),
	router.GET("/needs-logged-in-user", dummy_handler.Ping).
		Middlewares(
			application.Instance.Auth.GetProvider("session").GetAuthenticatorMiddleware(),
		),
}
