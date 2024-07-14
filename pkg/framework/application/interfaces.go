package application

import "github.com/pyramid.io/planit-backend/pkg/framework/session"

type TerminateableServiceInterface interface {
	Terminate()
}

// config interfaces which is required by application
type ConfigInterface interface {
	GetModulesConfig() []ModuleInterface
	GetSeverConfig() ServerConfigInterface
	GetLoggerConfig() LoggerConfigInterface
	GetSessionConfig() SessionConfigInterface
}

type ServerConfigInterface interface {
	GetPort() string
}

type LoggerConfigInterface interface {
	GetPath() string
}

type SessionConfigInterface interface {
	GetDriver() session.DriverKeyOrConstructor
	GetDriverConfig() map[string]interface{}
}

type ModuleInterface interface {
	Boot()
}