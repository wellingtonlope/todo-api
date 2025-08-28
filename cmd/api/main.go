package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
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
				fx.As(new(usecase.Clock)),
			),
			fx.Annotate(
				memory.NewTodoRepository,
				fx.As(new(todo.CreateStore)),
				fx.As(new(todo.GetAllStore)),
				fx.As(new(todo.GetByIDStore)),
				fx.As(new(todo.DeleteByIDStore)),
				fx.As(new(todo.UpdateStore)),
			),
			fx.Annotate(
				todo.NewCreate,
				fx.As(new(todo.Create)),
			),
			fx.Annotate(
				todo.NewGetAll,
				fx.As(new(todo.GetAll)),
			),
			fx.Annotate(
				todo.NewGetByID,
				fx.As(new(todo.GetByID)),
			),
			fx.Annotate(
				todo.NewDeleteByID,
				fx.As(new(todo.DeleteByID)),
			),
			fx.Annotate(
				todo.NewUpdate,
				fx.As(new(todo.Update)),
			),
			fx.Annotate(
				handler.NewTodoCreate,
				fx.As(new(handler.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				handler.NewTodoGetAll,
				fx.As(new(handler.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				handler.NewTodoGetByID,
				fx.As(new(handler.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				handler.NewTodoDeleteByID,
				fx.As(new(handler.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				handler.NewTodoUpdate,
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
