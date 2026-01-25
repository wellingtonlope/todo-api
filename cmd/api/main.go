// @title Todo API
// @version 1.0
// @description API for managing todo items
// @host localhost:1323
// @BasePath /
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wellingtonlope/todo-api/internal/bootstrap"
	"go.uber.org/fx"
)

func main() {
	if err := godotenv.Load(); err != nil && os.Getenv("APP_ENV") != "production" {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fx.New(
		bootstrap.FXOptions(),
	).Run()
}
