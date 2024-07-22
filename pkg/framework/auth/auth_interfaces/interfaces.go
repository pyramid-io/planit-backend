package auth_interfaces

import (
	"github.com/pyramid.io/planit-backend/pkg/framework/auth/user"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/middleware"
)

type AuthServiceInterface interface {
	AddProvider(name string, provider AuthProviderInterface) error
	GetProvider(name string) AuthProviderInterface
	User() (user.UserInterface, error)
	SetUser(user user.UserInterface)
}

type AuthProviderInterface interface {
	Login(username string, passsword string) error
	Logout(key string) error
	Authenticate(key string) (user.UserInterface, bool)
	GetAuthenticatorMiddleware() middleware.Middleware
}

type AuthProviderUserSourceInterface interface{}
