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

func (m *internalModule) Boot() {
	m.moduleRoutes = routes.GetRoutes()
	application.Instance.Router.RegisterRoutes(
		m.moduleRoutes,
	)
}