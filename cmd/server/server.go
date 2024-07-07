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
	fmt.Println("Application started.")

	if err != nil {
		log.Fatalf("unable to instantiate application instance")
	}

	defer app.Terminate()

	fmt.Println("Starting server...")
	app.StartServer()
	fmt.Println("Server started.")

}
