package routes

import (
	"fmt"
	"net/http"

	"github.com/pyramid.io/planit-backend/internal/handlers/dummy_handler"
	json_handler "github.com/pyramid.io/planit-backend/internal/handlers/json_handler"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
)

func GetRoutes() *[]router.Route {
	return &routes
}

var routes = []router.Route{
	{
		Path: "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "This is the HomePage")
		},
	},
	{
		Path:    "/ping",
		Handler: dummy_handler.Ping,
	},
	{
		Path:    "/pong",
		Handler: dummy_handler.Pong,
	},
	{
		Path:    "/json",
		Handler: json_handler.JsonResponseHandler,
	},
	{
		Path:    "/test-logger",
		Handler: dummy_handler.TestLogger,
	},
}
