package bootstrap

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	gormRepo "github.com/wellingtonlope/todo-api/internal/infra/gorm"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
	"github.com/wellingtonlope/todo-api/pkg/clock"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// provideMiddlewares returns the middleware functions used by both environments
func provideMiddlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		handler.Error,
	}
}

// provideEcho creates an Echo instance with optional lifecycle hooks
func provideEcho(config Config, middlewares []echo.MiddlewareFunc) *echo.Echo {
	e := echo.New()
	e.Use(middlewares...)

	return e
}

// provideEchoWithLifecycle creates an Echo instance with lifecycle hooks
func provideEchoWithLifecycle(config Config, middlewares []echo.MiddlewareFunc, lc fx.Lifecycle) *echo.Echo {
	e := echo.New()
	e.Use(middlewares...)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := e.Start(fmt.Sprintf(":%s", config.Port))
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
}

// provideDatabase creates a GORM database connection
func provideDatabase(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&gormRepo.TodoModel{}); err != nil {
		return nil, err
	}
	return db, nil
}

// provideClock returns the UTC clock with FX annotation
func provideClock() interface{} {
	return fx.Annotate(
		clock.NewClientUTC,
		fx.As(new(usecase.Clock)),
	)
}

// provideRepository returns the GORM todo repository with all interfaces
func provideRepository(db *gorm.DB) []interface{} {
	return []interface{}{
		fx.Annotate(
			gormRepo.NewTodoRepository,
			fx.As(new(todo.CreateStore)),
			fx.As(new(todo.GetAllStore)),
			fx.As(new(todo.GetByIDStore)),
			fx.As(new(todo.DeleteByIDStore)),
			fx.As(new(todo.UpdateStore)),
			fx.As(new(todo.CompleteStore)),
			fx.As(new(todo.MarkAsPendingStore)),
		),
	}
}

// provideUseCases returns all use case providers
func provideUseCases(
	createStore todo.CreateStore,
	getAllStore todo.GetAllStore,
	getByIDStore todo.GetByIDStore,
	deleteByIDStore todo.DeleteByIDStore,
	updateStore todo.UpdateStore,
	completeStore todo.CompleteStore,
	markPendingStore todo.MarkAsPendingStore,
	clock usecase.Clock,
) []interface{} {
	return []interface{}{
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
			todo.NewComplete,
			fx.As(new(todo.Complete)),
		),
		fx.Annotate(
			todo.NewMarkAsPending,
			fx.As(new(todo.MarkAsPending)),
		),
	}
}

// provideHandlers returns all handler providers
func provideHandlers(
	create todo.Create,
	getAll todo.GetAll,
	getByID todo.GetByID,
	deleteByID todo.DeleteByID,
	update todo.Update,
	complete todo.Complete,
	markPending todo.MarkAsPending,
) []interface{} {
	return []interface{}{
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
		fx.Annotate(
			handler.NewTodoComplete,
			fx.As(new(handler.Handler)),
			fx.ResultTags(`group:"handlers"`),
		),
		fx.Annotate(
			handler.NewTodoMarkPending,
			fx.As(new(handler.Handler)),
			fx.ResultTags(`group:"handlers"`),
		),
	}
}

// provideHandlerRegistration registers all handlers in Echo
func provideHandlerRegistration() interface{} {
	return fx.Annotate(
		func(handlers []handler.Handler, e *echo.Echo) {
			for _, h := range handlers {
				e.Add(h.Method(), h.Path(), h.Handle)
			}
		},
		fx.ParamTags(`group:"handlers"`),
	)
}

// provideSwaggerRegistration adds Swagger documentation
func provideSwaggerRegistration() interface{} {
	return func(e *echo.Echo) {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}
