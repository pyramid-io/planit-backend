package config

import (
	"github.com/pyramid.io/planit-backend/pkg/framework/application"
	"github.com/pyramid.io/planit-backend/pkg/framework/session"
)

type Config struct {
	Name     *string
	Modules  *[]application.ModuleInterface
	Server   application.ServerConfigInterface
	Logger   application.LoggerConfigInterface
	Session  application.SessionConfigInterface
	Services *map[string]interface{}
}

func (c *Config) GetModulesConfig() []application.ModuleInterface {
	return *c.Modules
}

func (c *Config) GetSeverConfig() application.ServerConfigInterface {
	return c.Server
}

func (c *Config) GetLoggerConfig() application.LoggerConfigInterface {
	return c.Logger
}

func (c *Config) GetSessionConfig() application.SessionConfigInterface {
	return c.Session
}

type ServerConfig struct {
	Port string
}

func (server *ServerConfig) GetPort() string {
	return server.Port
}

type LoggerConfig struct {
	Path string
}

func (logger *LoggerConfig) GetPath() string {
	return logger.Path
}

type SessionConfig struct {
	Driver session.DriverKeyOrConstructor
	Config map[string]interface{}
}

func (session *SessionConfig) GetDriver() session.DriverKeyOrConstructor {
	return session.Driver
}

func (session *SessionConfig) GetDriverConfig() map[string]interface{} {
	return session.Config
}
