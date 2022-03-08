package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/infra/http"
	"github.com/wellingtonlope/todo-api/internal/infra/mongo"
)

func main() {
	if err := godotenv.Load(); err != nil && os.Getenv("APP_ENV") != "production" {
		log.Fatalf("Error loading .env file: %v", err)
	}

	repositories := &mongo.Repositories{
		UriConnection: os.Getenv("MONGO_URI"),
		Database:      os.Getenv("MONGO_DATABASE"),
	}

	useCases, err := usecase.GetUseCases(repositories)
	if err != nil {
		log.Fatalf("Error during server initialization: %v", err)
	}

	e := echo.New()
	http.InitHandlers(e, useCases)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
