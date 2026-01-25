package bootstrap

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	gormRepo "github.com/wellingtonlope/todo-api/internal/infra/gorm"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
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
