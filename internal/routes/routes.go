package routes

import (
	"fmt"
	"net/http"

	"github.com/pyramid.io/planit-backend/internal/handlers/dummyHandler"
	jsonhandler "github.com/pyramid.io/planit-backend/internal/handlers/jsonHandler"
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
		Handler: dummyHandler.Ping,
	},
	{
		Path:    "/pong",
		Handler: dummyHandler.Pong,
	},
	{
		Path:    "/json",
		Handler: jsonhandler.JsonResponseHandler,
	},
	{
		Path:    "/test-logger",
		Handler: dummyHandler.TestLogger,
	},
}
