package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/pyramid.io/planit-backend/config"
	"github.com/pyramid.io/planit-backend/pkg/framework/application"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println("Starting the application...")
	app, err := application.New(config.GetInstance())
	defer app.Terminate()

	if err != nil {
		log.Fatalf("unable to instantiate application instance")
	}

	fmt.Println("Starting server...")
	app.StartServer()
}
