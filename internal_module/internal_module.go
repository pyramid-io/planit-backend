package internal_module

import (
	"github.com/pyramid.io/planit-backend/internal_module/routes"
	"github.com/pyramid.io/planit-backend/pkg/framework/application"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
)

var Module application.ModuleInterface = &internalModule{}

type internalModule struct {
	moduleRoutes *[]router.RouteInterface
}

func (m *internalModule) Boot(application *application.Application) {
	m.moduleRoutes = routes.GetRoutes()
	application.Router.RegisterRoutes(
		m.moduleRoutes,
	)
}