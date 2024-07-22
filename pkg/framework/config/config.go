package config

import (
	"github.com/pyramid.io/planit-backend/pkg/framework/application/interfaces"
)

type Config struct {
	Name     *string
	Modules  *[]interfaces.ModuleInterface
	Server   interfaces.ServerConfigInterface
	Logger   interfaces.LoggerConfigInterface
	Session  interfaces.SessionConfigInterface
	Database []interfaces.DatabaseDriverConfigInterface
	Auth []interfaces.AuthenticationProviderConfigInterface
	Services *map[string]interface{}
}

func (c *Config) GetModulesConfig() []interfaces.ModuleInterface {
	return *c.Modules
}

func (c *Config) GetSeverConfig() interfaces.ServerConfigInterface {
	return c.Server
}

func (c *Config) GetLoggerConfig() interfaces.LoggerConfigInterface {
	return c.Logger
}

func (c *Config) GetSessionConfig() interfaces.SessionConfigInterface {
	return c.Session
}

func (c *Config) GetDatabaseConfig() []interfaces.DatabaseDriverConfigInterface {
	return c.Database
}

func (c *Config) GetAuthConfig() []interfaces.AuthenticationProviderConfigInterface {
	return c.Auth
}

//------------------------------------------------------------//
type ServerConfig struct {
	Port string
}

func (server *ServerConfig) GetPort() string {
	return server.Port
}

//------------------------------------------------------------//
type LoggerConfig struct {
	Path string
}

func (logger *LoggerConfig) GetPath() string {
	return logger.Path
}

//------------------------------------------------------------//
type SessionConfig struct {
	Driver interfaces.DriverKeyOrConstructor
	Config map[string]interface{}
}

func (session *SessionConfig) GetDriver() interfaces.DriverKeyOrConstructor {
	return session.Driver
}

func (session *SessionConfig) GetDriverConfig() map[string]interface{} {
	return session.Config
}

//------------------------------------------------------------//
type DatabaseConnectionConfig struct {
	Driver         interfaces.DriverKeyOrConstructor
	ConnectionName string
	Config         map[string]interface{}
}

func (d DatabaseConnectionConfig) GetDriver() interfaces.DriverKeyOrConstructor {
	return d.Driver
}

func (d DatabaseConnectionConfig) GetConnectionName() string {
	return d.ConnectionName
}

func (d DatabaseConnectionConfig) GetConfig() map[string]interface{} {
	return d.Config
}

//------------------------------------------------------------//
type AuthenticationProviderConfig struct {
	Provider interfaces.DriverKeyOrConstructor
	Name string
	Config map[string]interface{}
}

func (d AuthenticationProviderConfig) GetProvider() interfaces.DriverKeyOrConstructor {
	return d.Provider
}

func (d AuthenticationProviderConfig) GetName() string {
	return d.Name
}

func (d AuthenticationProviderConfig) GetConfig() map[string]interface{} {
	return d.Config
}
