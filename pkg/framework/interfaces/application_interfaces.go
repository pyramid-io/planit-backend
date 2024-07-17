package interfaces

type TerminateableServiceInterface interface {
	Terminate()
}

type ModuleInterface interface {
	Boot()
	GetName() string
	GetResourceDir() string
}