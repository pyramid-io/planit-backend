package session

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/pyramid.io/planit-backend/pkg/framework/interfaces"
)

var serviceRegistry = map[string]DriverConstructor{
	"filesystem": NewFileSystemSessionDriver,
}

type DriverConstructor func(config map[string]interface{}) (SessionDriverInterface, error)

type SessionServiceInterface interface {
	Create(data map[string]interface{}, expiry *time.Time) (*Session, error)
	Get(id string) (*Session, error)
	Delete(id string) error
}

type SessionDriverInterface interface {
	Create(key string, data map[string]interface{}, expiresAt *time.Time) (*Session, error)
	Get(key string) (*Session, error)
	Delete(key string) error
}

type Session struct {
	ID string
	Data map[string]interface{}
	ExpiresAt *time.Time
}

type SessionService struct {
	driver SessionDriverInterface
}

func New(driverKeyOrConstructor interfaces.DriverKeyOrConstructor, driverConfig map[string]interface{}) (SessionServiceInterface, error) {
	var constructor DriverConstructor
	var ok bool

	switch v := driverKeyOrConstructor.(type) {
	case string:
		constructor, ok = serviceRegistry[v]
		if !ok {
			return nil, errors.New("session driver could not be started")
		}
	case DriverConstructor:
		constructor = v
	default:
		return nil, errors.New("invalid driver key or constructor")
	}

	driver, err := constructor(driverConfig)
	if err != nil {
		return nil, err
	}

	return &SessionService{
		driver: driver,
	}, nil
}

func (service *SessionService) Create(data map[string]interface{}, expiry *time.Time) (*Session, error) {
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

func (service *SessionService) Get(id string) (*Session, error) {
	session, err := service.driver.Get(id)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (service *SessionService) Delete(id string) error {
	return service.driver.Delete(id)
}
