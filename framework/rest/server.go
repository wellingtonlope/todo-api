package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/application/usecase"
	"github.com/wellingtonlope/todo-api/framework/db/mongodb"
	"github.com/wellingtonlope/todo-api/framework/rest/handler"
)

func main() {
	if err := godotenv.Load(); err != nil && os.Getenv("APP_ENV") != "production" {
		log.Fatalf("Error loading .env file: %v", err)
	}

	repositories := mongodb.Repositories{
		UriConnection: os.Getenv("MONGO_URI"),
		Database:      os.Getenv("MONGO_DATABASE"),
	}

	useCases, err := usecase.GetUseCases(&repositories)
	if err != nil {
		log.Fatalf("Error during server initialization: %v", err)
	}

	e := echo.New()
	handler.InitHandlers(e, useCases)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
