package interfaces

type ApplicationInterface interface{
	Boot()
	StartServer()
	Terminate() 
	GetModule(key string) (ModuleInterface, error)

}
type TerminateableServiceInterface interface {
	Terminate()
}

type ModuleInterface interface {
	Boot()
	GetName() string
	GetResourceDir() string
}