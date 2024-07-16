package interfaces

type TerminateableServiceInterface interface {
	Terminate()
}

type ModuleInterface interface {
	Boot()
}