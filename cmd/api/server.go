package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil && os.Getenv("APP_ENV") != "production" {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fx.New(
		fx.Provide(
			func(lc fx.Lifecycle) *echo.Echo {
				e := echo.New()
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
							if err != nil {
								e.Logger.Fatal(err)
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return e.Shutdown(ctx)
					},
				})
				return e
			},
			func() (*gorm.DB, error) {
				return gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
			},
		),
	).Run()
}
