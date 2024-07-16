package interfaces

type ConfigInterface interface {
	GetModulesConfig() []ModuleInterface
	GetSeverConfig() ServerConfigInterface
	GetLoggerConfig() LoggerConfigInterface
	GetSessionConfig() SessionConfigInterface
	GetDatabaseConfig() []DatabaseConnectionConfigInterface
}

type ServerConfigInterface interface {
	GetPort() string
}

type LoggerConfigInterface interface {
	GetPath() string
}

type DriverKeyOrConstructor interface{}

type SessionConfigInterface interface {
	GetDriver() DriverKeyOrConstructor
	GetDriverConfig() map[string]interface{}
}

type DatabaseConfigInterface interface {}

type DatabaseConnectionConfigInterface interface {
	GetDriver() DriverKeyOrConstructor
	GetConnectionName() string
	GetConfig() map[string]interface{}
}