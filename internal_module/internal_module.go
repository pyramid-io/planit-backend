package internal_module

import (
	"path/filepath"

	"github.com/pyramid.io/planit-backend/internal_module/routes"
	"github.com/pyramid.io/planit-backend/pkg/framework/application"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
	"github.com/pyramid.io/planit-backend/pkg/framework/application/interfaces"
)

var Module interfaces.ModuleInterface = &internalModule{
	Name: "internal",
}

type internalModule struct {
	Name string
	moduleRoutes *[]router.RouteInterface
}

func (m *internalModule) GetName() string{
	return m.Name
}

func (m *internalModule) Boot() {
	m.moduleRoutes = routes.GetRoutes()
	application.Instance.Router.RegisterRoutes(
		m.moduleRoutes,
	)
}

func (m *internalModule) GetResourceDir() string {
	return filepath.Join(
		*application.Instance.Dir,
		"internal_module",
		"resources",
	)
}