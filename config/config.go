package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/pyramid.io/planit-backend/internal_module"
	"github.com/pyramid.io/planit-backend/pkg/framework/application/interfaces"
	"github.com/pyramid.io/planit-backend/pkg/framework/config"
	"github.com/pyramid.io/planit-backend/pkg/framework/utils"
)


var (
	instance *config.Config
	once     sync.Once
)

func GetInstance() *config.Config {
	once.Do(initialize)
	return instance
}

func initialize() {
	fmt.Println("Initializing config...")

	appName := utils.ReadEnv("APP_NAME", "planit")

	instance = &config.Config{
		Name: &appName,
		Modules: &[]interfaces.ModuleInterface{
			internal_module.Module,
		},
		Server: &config.ServerConfig{
			Port: utils.ReadEnvOrPanic("SERVER_PORT"),
		},
		Logger: &config.LoggerConfig{
			Path: utils.ReadEnvOrPanic("LOG_PATH"),
		},
		Session: &config.SessionConfig{
			Driver: "filesystem",
			Config: map[string]interface{}{
				"dir":        utils.ReadEnvOrPanic("SESSION_PATH"),
				"defaultTTL": 30 * time.Minute,
			},
		},
		Database: []interfaces.DatabaseDriverConfigInterface{
			config.DatabaseConnectionConfig{
				Driver: "mysql",
				ConnectionName: "api",
				Config: map[string]interface{}{
					"username": utils.ReadEnvOrPanic("MYSQL_API_USERNAME"),
					"password": utils.ReadEnvOrPanic("MYSQL_API_PASSWORD"),
					"host": utils.ReadEnvOrPanic("MYSQL_API_HOST"),
					"port": utils.ReadEnv("MYSQL_API_PORT", "3306"),
					"databaseName": utils.ReadEnvOrPanic("MYSQL_API_DATABASE"),
					"charset": "utf8mb4",
				},
			},
		},
		Auth: []interfaces.AuthenticationProviderConfigInterface{
			config.AuthenticationProviderConfig{
				Provider: "session",
				Name: "session",
				Config: make(map[string]interface{}),
			},
		},
	}

	fmt.Println("Config initialized...")
}
