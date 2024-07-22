package drivers

import "github.com/pyramid.io/planit-backend/pkg/framework/session/session_interfaces"

type SessionDriverConstructor func(config map[string]interface{}) (session_interfaces.SessionDriverInterface, error)

var ServiceRegistry = map[string]SessionDriverConstructor{
	"filesystem": NewFileSystemSessionDriver,
}