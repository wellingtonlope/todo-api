package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
	"github.com/wellingtonlope/todo-api/internal/infra/memory"
	"github.com/wellingtonlope/todo-api/pkg/clock"
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
			func() []echo.MiddlewareFunc {
				return []echo.MiddlewareFunc{
					handler.Error,
				}
			},
			func(lc fx.Lifecycle, middlewares []echo.MiddlewareFunc) *echo.Echo {
				e := echo.New()
				e.Use(middlewares...)
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
			fx.Annotate(
				clock.NewClientUTC,
				fx.As(new(clock.Client)),
			),
			fx.Annotate(
				memory.NewTodoRepository,
				fx.As(new(todo.CreateStore)),
				fx.As(new(todo.GetByIDStore)),
			),
			fx.Annotate(
				todo.NewCreate,
				fx.As(new(todo.Create)),
			),
			fx.Annotate(
				todo.NewGetByID,
				fx.As(new(todo.GetByID)),
			),
			fx.Annotate(
				handler.NewTodoCreate,
				fx.As(new(handler.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				handler.NewTodoGetByID,
				fx.As(new(handler.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
		),
		fx.Invoke(
			fx.Annotate(
				func(handlers []handler.Handler, e *echo.Echo) {
					for _, h := range handlers {
						e.Add(h.Method(), h.Path(), h.Handle)
					}
				},
				fx.ParamTags(`group:"handlers"`),
			),
		),
	).Run()
}
