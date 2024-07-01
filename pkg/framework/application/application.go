package application

import (
	"reflect"

	"github.com/knadh/koanf"
	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
	"github.com/pyramid.io/planit-backend/pkg/framework/logger"
	"github.com/pyramid.io/planit-backend/pkg/framework/server"
)

type Application struct {
	Config *koanf.Koanf
	Router router.RouterInterface
	Server server.ServerInterface
	Logger logger.LoggerInterface
}

var application *Application

func GetInstance() *Application {
	return application
}

func New(config *koanf.Koanf, routes *[]router.Route) (*Application, error) {
	router := router.New()
	router.RegisterRoutes(routes)

	server := server.New(router)

	logger, err := logger.New(config.String("logger.path"))
	if err != nil {
		return nil, err
	}

	application = &Application{
		Config: config,
		Router: router,
		Server: server,
		Logger: logger,
	}

	return application, nil
}

func (application *Application) Start(port string) {
	application.Server.Start(application.Config.String("server.port"))
}

func (application *Application) Terminate() {

	v := reflect.ValueOf(application)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		iface, ok := field.Interface().(TerminateableServiceInterface)
		if ok {
			iface.Terminate()
		}
	}
}
