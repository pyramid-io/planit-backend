package application

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
	"github.com/pyramid.io/planit-backend/pkg/framework/logger"
	"github.com/pyramid.io/planit-backend/pkg/framework/server"
)

type Application struct {
	Config ConfigInterface
	Modules []ModuleInterface
	Router router.RouterInterface
	Server server.ServerInterface
	Logger logger.LoggerInterface
}

var (
	Instance *Application
	once sync.Once
)

func New(config ConfigInterface) (*Application, error) {
	fmt.Println("Initializing framework application...")
	
	modules := config.GetModulesConfig()
	router, err := router.New()
	if err != nil {
		return nil, err
	}

	server, err := server.New(router)
	if err != nil {
		return nil, err
	}

	logger, err := logger.New(config.GetLoggerConfig().GetPath())
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		Instance = &Application{
			Config: config,
			Modules: modules,
			Router: router,
			Server: server,
			Logger: logger,
		}
	})

	Instance.Boot()
	
	return Instance, nil
}

func (appInstance *Application) Boot() {
	for _, module := range appInstance.Modules {
        module.Boot(appInstance)
    }
}

func (appInstance *Application) StartServer() {
	appInstance.Server.Start(
		appInstance.Config.GetSeverConfig().GetPort(),
	)
}

func (appInstance *Application) Terminate() {

	v := reflect.ValueOf(appInstance)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		iface, ok := field.Interface().(TerminateableServiceInterface)
		if ok {
			iface.Terminate()
		}
	}
}
