package session

import (
	"errors"
	"time"

	"github.com/google/uuid"
	application_interfaces "github.com/pyramid.io/planit-backend/pkg/framework/application/interfaces"
	"github.com/pyramid.io/planit-backend/pkg/framework/session/drivers"
	"github.com/pyramid.io/planit-backend/pkg/framework/session/session_interfaces"
)
type SessionManager struct {
	driver session_interfaces.SessionDriverInterface
}

func New(driverKeyOrConstructor application_interfaces.DriverKeyOrConstructor, driverConfig map[string]interface{}) (session_interfaces.SessionManagerInterface, error) {
	var constructor drivers.SessionDriverConstructor
	var ok bool

	switch v := driverKeyOrConstructor.(type) {
	case string:
		constructor, ok = drivers.ServiceRegistry[v]
		if !ok {
			return nil, errors.New("session driver could not be started")
		}
	case drivers.SessionDriverConstructor:
		constructor = v
	default:
		return nil, errors.New("invalid driver key or constructor")
	}

	driver, err := constructor(driverConfig)
	if err != nil {
		return nil, err
	}

	return &SessionManager{
		driver: driver,
	}, nil
}

func (service *SessionManager) Create(data map[string]interface{}, expiry *time.Time) (session_interfaces.SessionInterface, error) {
	sessionID := uuid.New().String()

	session, err := service.driver.Create(
		sessionID,
		data,
		expiry,
	)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (service *SessionManager) Get(id string) (session_interfaces.SessionInterface, error) {
	session, err := service.driver.Get(id)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (service *SessionManager) Delete(id string) error {
	return service.driver.Delete(id)
}
