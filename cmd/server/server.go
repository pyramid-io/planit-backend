package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pyramid.io/planit-backend/internal/routes"
	"github.com/pyramid.io/planit-backend/pkg/configloader"
	"github.com/pyramid.io/planit-backend/pkg/framework/application"
)

func main() {
	cfgPath := getConfigPath()
	cfg := configloader.Load(cfgPath)

	app, err := application.New(cfg, routes.GetRoutes())
	if err != nil {
		log.Fatalf("unable to instantiate application instance")
	}
	defer app.Terminate()

	app.Start(cfg.String("server.port"))
}

func getConfigPath() string {
	err := godotenv.Load("/var/www/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("unable to run the app, CONFIG_PATH env is missing.")
	}

	return configPath
}
