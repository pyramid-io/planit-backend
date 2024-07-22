package providers

import (
	"net/http"

	"github.com/pyramid.io/planit-backend/pkg/framework/auth/auth_interfaces"
	"github.com/pyramid.io/planit-backend/pkg/framework/auth/user"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/middleware"
	"github.com/pyramid.io/planit-backend/pkg/framework/session/session_interfaces"
)

func New(userSource auth_interfaces.AuthProviderUserSourceInterface) *SessionAuthProvider {
	return &SessionAuthProvider{
		userSource: userSource,
	}
}

type SessionAuthProvider struct {
	authService auth_interfaces.AuthServiceInterface
	userSource auth_interfaces.AuthProviderUserSourceInterface
	sessionManager session_interfaces.SessionManagerInterface
}

func (provider *SessionAuthProvider) Login(username string, passsword string) error {
	return nil
}

func (provider *SessionAuthProvider) Logout(key string) error {
	return nil
}

func (provider *SessionAuthProvider) Authenticate(key string) (user.UserInterface, bool) {
	return nil, false
}

func (provider *SessionAuthProvider) GetAuthenticatorMiddleware() middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userSessionValue := r.Header.Get("usersession")
            if userSessionValue == "" {
                http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
                return
            }

            user, isAuthenticated := provider.Authenticate(userSessionValue)
            if !isAuthenticated {
                http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
                return
            }

			provider.authService.SetUser(user)

            next.ServeHTTP(w, r)
		})
	}
}


func SessionAuthProviderConstructor(
	authService auth_interfaces.AuthServiceInterface,
	config map[string]interface{},
) (auth_interfaces.AuthProviderInterface, error) {
	return nil, nil
}