package main

import (
	"fmt"
	"log"
	"os"

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
	app, err := application.New(config.GetInstance(), getRootDir())
	defer app.Terminate()

	if err != nil {
		log.Fatalf("unable to instantiate application instance: ", err)
	}

	fmt.Println("Starting server...")
	app.StartServer()
}

func getRootDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	return wd
}
