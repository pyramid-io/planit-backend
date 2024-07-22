package session_interfaces

import (
	"time"
)

type SessionInterface interface {
	GetID() string
	GetData() map[string]interface{}
	GetExpiresAt() *time.Time
}

type SessionManagerInterface interface {
	Create(data map[string]interface{}, expiry *time.Time) (SessionInterface, error)
	Get(id string) (SessionInterface, error)
	Delete(id string) error
}

type SessionDriverInterface interface {
	Create(key string, data map[string]interface{}, expiresAt *time.Time) (SessionInterface, error)
	Get(key string) (SessionInterface, error)
	Delete(key string) error
}