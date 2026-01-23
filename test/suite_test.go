package test

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
	"github.com/wellingtonlope/todo-api/internal/infra/memory"
	"github.com/wellingtonlope/todo-api/pkg/clock"
	"github.com/wellingtonlope/todo-api/test/steps"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTodoCreationBDD(t *testing.T) {
	// Setup DB
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		t.Fatal(err)
	}

	// Setup dependencies
	clock := clock.NewClientUTC()
	store := memory.NewTodoRepository()
	createUsecase := todo.NewCreate(store, clock)
	createHandler := handler.NewTodoCreate(createUsecase)

	// Setup Echo
	e := echo.New()
	e.Use(handler.Error) // Add error handling middleware
	e.POST("/todos", createHandler.Handle)

	tc := &steps.TodoCreationContext{
		EchoApp: e,
		DB:      db,
	}

	suite := godog.TestSuite{
		ScenarioInitializer: tc.InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/todo_creation.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func TestTodoGetByIDBDD(t *testing.T) {
	// Setup DB
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		t.Fatal(err)
	}

	// Setup dependencies
	clock := clock.NewClientUTC()
	store := memory.NewTodoRepository()
	createUsecase := todo.NewCreate(store, clock)
	createHandler := handler.NewTodoCreate(createUsecase)
	getByIDUsecase := todo.NewGetByID(store)
	getByIDHandler := handler.NewTodoGetByID(getByIDUsecase)

	// Setup Echo
	e := echo.New()
	e.Use(handler.Error) // Add error handling middleware
	e.POST("/todos", createHandler.Handle)
	e.GET("/todos/:id", getByIDHandler.Handle)

	tc := &steps.TodoGetByIDContext{
		EchoApp: e,
		DB:      db,
	}

	suite := godog.TestSuite{
		ScenarioInitializer: tc.InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/todo_get_by_id.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
