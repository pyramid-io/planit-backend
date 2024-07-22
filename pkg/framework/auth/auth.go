package auth

import (
	"errors"

	"github.com/pyramid.io/planit-backend/pkg/framework/auth/auth_interfaces"
	"github.com/pyramid.io/planit-backend/pkg/framework/auth/user"
	app_interfaces "github.com/pyramid.io/planit-backend/pkg/framework/application/interfaces"
)

type AuthService struct {
	appInstance app_interfaces.ApplicationInterface
	providers   map[string]interfaces.AuthProviderInterface
	user        user.UserInterface
}

func (authService *AuthService) AddProvider(name string, provider interfaces.AuthProviderInterface) error {
	authService.providers[name] = provider

	return nil
}

func (authService *AuthService) GetProvider(name string) interfaces.AuthProviderInterface {
	provider, exists := authService.providers[name]
	if !exists {
		return  nil
	}

	return provider
}

func (authService *AuthService) User() (user.UserInterface, error) {
	if authService.user != nil {
		return authService.user, nil
	}

	return nil, errors.New("no user is authenticated on the auth service.")
}

func (authService *AuthService) SetUser(user user.UserInterface) {
	authService.user = user
}

var authService interfaces.AuthServiceInterface = &AuthService{}

func New(appInstance *app_interfaces.ApplicationInterface, providersConfig []app_interfaces.AuthenticationProviderConfigInterface) (interfaces.AuthServiceInterface, error) {
	authService = appInstance.Session.
	for _, providerConfig := range providersConfig {
		var constructor AuthenticationProviderConstructor
		var ok bool

		switch v := providerConfig.GetDriver().(type) {
		case string:
			constructor, ok = ServiceRegistry[v]
			if !ok {
				return nil, errors.New("auth provider by name is not found in registry")
			}
		case AuthenticationProviderConstructor:
			constructor = v
		default:
			return nil, errors.New("invalid connection key or constructor")
		}

		authProvider, err := constructor(authService, providerConfig.GetConfig())
		if err != nil {
			return nil, err
		}

		authService.AddProvider(
			providerConfig.GetName(),
			authProvider.GetConfig(),
		)
	}

	return authService, nil
}
