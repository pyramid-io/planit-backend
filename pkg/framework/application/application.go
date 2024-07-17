package application

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/pyramid.io/planit-backend/pkg/framework/database"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
	"github.com/pyramid.io/planit-backend/pkg/framework/interfaces"
	"github.com/pyramid.io/planit-backend/pkg/framework/logger"
	"github.com/pyramid.io/planit-backend/pkg/framework/server"
	"github.com/pyramid.io/planit-backend/pkg/framework/session"
)

type Application struct {
	Dir *string
	Config  interfaces.ConfigInterface
	Modules map[string]interfaces.ModuleInterface
	Router  router.RouterInterface
	Server  server.ServerInterface
	Logger  logger.LoggerInterface
	Session session.SessionServiceInterface
	Database database.DatabaseServiceInterface
}

var (
	Instance *Application
	once     sync.Once
)

func New(config interfaces.ConfigInterface, dir string) (*Application, error) {
	fmt.Println("Initializing framework application...")

	modulesConfig := config.GetModulesConfig()

	modules := make(map[string]interfaces.ModuleInterface)
	for _, module := range modulesConfig {
		modules[module.GetName()] = module 
	}


	
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

	session, err := session.New(
		config.GetSessionConfig().GetDriver(),
		config.GetSessionConfig().GetDriverConfig(),
	)

	if err != nil {
		return nil, err
	}

	database, err := database.New(config.GetDatabaseConfig())
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		Instance = &Application{
			Dir:      &dir,
			Config:   config,
			Modules:  modules,
			Router:   router,
			Server:   server,
			Logger:   logger,
			Session:  session,
			Database: database,
		}
	})

	Instance.Boot()

	return Instance, nil
}

func (appInstance *Application) Boot() {
	for _, module := range appInstance.Modules {
		module.Boot()
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
		iface, ok := field.Interface().(interfaces.TerminateableServiceInterface)
		if ok {
			iface.Terminate()
		}
	}
}

func (appInstance *Application) GetModule(key string) (interfaces.ModuleInterface, error) {
	if module, exists := appInstance.Modules[key]; exists {
		return module, nil
	}

	return nil, errors.New("module is not registered for key: " + key)
}
