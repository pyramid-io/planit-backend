package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/pyramid.io/planit-backend/internal_module"
	"github.com/pyramid.io/planit-backend/pkg/framework/application"
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
		Modules: &[]application.ModuleInterface{
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
				"dir": utils.ReadEnvOrPanic("SESSION_PATH"),
				"defaultTTL": 30 * time.Minute,
			},
		},
	}

	fmt.Println("Config initialized...")
}
