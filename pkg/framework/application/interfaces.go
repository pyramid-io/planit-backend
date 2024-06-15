package application

type TerminateableServiceInterface interface {
	Terminate()
}

type ModuleInterface interface {
	Boot(application *Application)
}


// config interfaces which is required by application
type ConfigInterface interface {
	GetModulesConfig() []ModuleInterface
	GetSeverConfig() ServerConfigInterface
	GetLoggerConfig() LoggerConfigInterface
}

type ServerConfigInterface interface {
	GetPort() string
}

type LoggerConfigInterface interface {
	GetPath() string
}
