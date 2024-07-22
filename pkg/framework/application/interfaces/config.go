package interfaces

type ConfigInterface interface {
	GetModulesConfig() []ModuleInterface
	GetSeverConfig() ServerConfigInterface
	GetLoggerConfig() LoggerConfigInterface
	GetSessionConfig() SessionConfigInterface
	GetDatabaseConfig() []DatabaseDriverConfigInterface
	GetAuthConfig() []AuthenticationProviderConfigInterface
}

type DriverKeyOrConstructor interface{}

type ServerConfigInterface interface {
	GetPort() string
}

type LoggerConfigInterface interface {
	GetPath() string
}

type SessionConfigInterface interface {
	GetDriver() DriverKeyOrConstructor
	GetDriverConfig() map[string]interface{}
}

type DatabaseDriverConfigInterface interface {
	GetDriver() DriverKeyOrConstructor
	GetConnectionName() string
	GetConfig() map[string]interface{}
}

type AuthenticationProviderConfigInterface interface{
	GetProvider() DriverKeyOrConstructor
	GetName() string
	GetConfig() map[string]interface{}
}
