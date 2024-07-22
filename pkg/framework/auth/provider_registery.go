package auth

import (
	"github.com/pyramid.io/planit-backend/pkg/framework/auth/auth_interfaces"
	"github.com/pyramid.io/planit-backend/pkg/framework/auth/providers"
)

type AuthenticationProviderConstructor func(
	authService auth_interfaces.AuthProviderInterface,
	config map[string]interface{},
) (auth_interfaces.AuthProviderInterface, error)

var ServiceRegistry = map[string]AuthenticationProviderConstructor{
	"session": providers.SessionAuthProviderConstructor,
}
