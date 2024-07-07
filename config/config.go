package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/pyramid.io/planit-backend/internal_module"
	"github.com/pyramid.io/planit-backend/pkg/framework/application"
	"github.com/pyramid.io/planit-backend/pkg/framework/utils"
)

var (
	instance *AppConfig
	once sync.Once
)

func GetInstance() *AppConfig {
	once.Do(initialize)
	return instance
}

func initialize() {
	fmt.Println("Initializing config...")

	appName := utils.ReadEnv("APP_NAME", "planit")

	instance = &AppConfig{
		name: &appName,
		modules: &[]application.ModuleInterface{
			internal_module.Module,
		},
		server: &ServerConfig{
			port: utils.ReadEnvOrPanic("SERVER_PORT"),
		},
		logger: &LoggerConfig{
			path: utils.ReadEnvOrPanic("LOG_PATH"),
		},
	}

	fmt.Println("Config initialized...")
}

type AppConfig struct {
	name *string
	modules *[]application.ModuleInterface
	server  *ServerConfig
	logger  *LoggerConfig
}

func (c *AppConfig) GetModulesConfig() []application.ModuleInterface {
	return *c.modules
}

func (c *AppConfig) GetSeverConfig() application.ServerConfigInterface {
	return c.server
}

func (c *AppConfig) GetLoggerConfig() application.LoggerConfigInterface {
	return c.logger
}

func (c *AppConfig) Get(path string) (interface{}, error) {
	fields := strings.Split(path, ".")
	var current reflect.Value = reflect.ValueOf(c).Elem()

	for _, field := range fields {
		current = current.FieldByName(field)
		if !current.IsValid() {
			return nil, errors.New("filed is not found")
		}
	}

	return current.Interface(), nil
}

type ServerConfig struct {
	port string
}

func (server *ServerConfig) GetPort() string {
	return server.port
}

type LoggerConfig struct {
	path string
}

func (logger *LoggerConfig) GetPath() string {
	return logger.path
}
