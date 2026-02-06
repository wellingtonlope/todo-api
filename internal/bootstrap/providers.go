package bootstrap

import (
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	gormRepo "github.com/wellingtonlope/todo-api/internal/infra/gorm"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
	"github.com/wellingtonlope/todo-api/pkg/clock"
	"go.uber.org/fx"
)

// CommonProviders returns fx.Option with all shared providers for clock, repositories, use cases, and handlers
func CommonProviders() fx.Option {
	providers := []interface{}{
		// Clock provider
		fx.Annotate(
			clock.NewClientUTC,
			fx.As(new(usecase.Clock)),
		),
		// Repository provider
		fx.Annotate(
			gormRepo.NewTodoRepository,
			fx.As(new(todo.CreateStore)),
			fx.As(new(todo.ListStore)),
			fx.As(new(todo.GetByIDStore)),
			fx.As(new(todo.DeleteByIDStore)),
			fx.As(new(todo.UpdateStore)),
			fx.As(new(todo.CompleteStore)),
			fx.As(new(todo.MarkAsPendingStore)),
		),
		// Use case providers
		fx.Annotate(
			todo.NewCreate,
			fx.As(new(todo.Create)),
		),
		fx.Annotate(
			todo.NewList,
			fx.As(new(todo.List)),
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
		// Handler providers
		fx.Annotate(
			handler.NewTodoCreate,
			fx.As(new(handler.Handler)),
			fx.ResultTags(`group:"handlers"`),
		),
		fx.Annotate(
			handler.NewTodoList,
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

	return fx.Module("common", fx.Provide(providers...))
}
