package test

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
	"github.com/wellingtonlope/todo-api/internal/infra/memory"
	"github.com/wellingtonlope/todo-api/pkg/clock"
	"github.com/wellingtonlope/todo-api/test/steps"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestDependencies struct {
	DB                 *gorm.DB
	Clock              usecase.Clock
	Store              interface{}
	CreateUsecase      todo.Create
	GetByIDUsecase     todo.GetByID
	DeleteByIDUsecase  todo.DeleteByID
	MarkPendingUsecase todo.MarkAsPending
	CreateHandler      *handler.TodoCreate
	GetByIDHandler     *handler.TodoGetByID
	DeleteByIDHandler  *handler.TodoDeleteByID
	MarkPendingHandler *handler.TodoMarkPending
}

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func setupDependencies(db *gorm.DB) *TestDependencies {
	clock := clock.NewClientUTC()
	store := memory.NewTodoRepository()
	createUsecase := todo.NewCreate(store, clock)
	getByIDUsecase := todo.NewGetByID(store)
	deleteByIDUsecase := todo.NewDeleteByID(store)
	markPendingUsecase := todo.NewMarkAsPending(store, clock)
	createHandler := handler.NewTodoCreate(createUsecase)
	getByIDHandler := handler.NewTodoGetByID(getByIDUsecase)
	deleteByIDHandler := handler.NewTodoDeleteByID(deleteByIDUsecase)
	markPendingHandler := handler.NewTodoMarkPending(markPendingUsecase)

	return &TestDependencies{
		DB:                 db,
		Clock:              clock,
		Store:              store,
		CreateUsecase:      createUsecase,
		GetByIDUsecase:     getByIDUsecase,
		DeleteByIDUsecase:  deleteByIDUsecase,
		MarkPendingUsecase: markPendingUsecase,
		CreateHandler:      createHandler,
		GetByIDHandler:     getByIDHandler,
		DeleteByIDHandler:  deleteByIDHandler,
		MarkPendingHandler: markPendingHandler,
	}
}

func setupEchoApp(deps *TestDependencies, includeGetByID bool) *echo.Echo {
	e := echo.New()
	e.Use(handler.Error)
	e.POST("/todos", deps.CreateHandler.Handle)
	if includeGetByID {
		e.GET("/todos/:id", deps.GetByIDHandler.Handle)
	}
	e.DELETE("/todos/:id", deps.DeleteByIDHandler.Handle)
	e.POST("/todos/:id/pending", deps.MarkPendingHandler.Handle)
	return e
}

func runBDDTest(t *testing.T, app *echo.Echo, db *gorm.DB, featurePaths []string, initializer func(*godog.ScenarioContext)) {
	suite := godog.TestSuite{
		ScenarioInitializer: initializer,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    featurePaths,
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func TestTodoCreationBDD(t *testing.T) {
	db := setupDatabase(t)
	deps := setupDependencies(db)
	app := setupEchoApp(deps, false) // false for no GetByID

	tc := &steps.TodoCreationContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_creation.feature"}, tc.InitializeScenario)
}

func TestTodoGetByIDBDD(t *testing.T) {
	db := setupDatabase(t)
	deps := setupDependencies(db)
	app := setupEchoApp(deps, true) // true for include GetByID

	tc := &steps.TodoGetByIDContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_get_by_id.feature"}, tc.InitializeScenario)
}

func TestTodoDeleteByIDBDD(t *testing.T) {
	db := setupDatabase(t)
	deps := setupDependencies(db)
	app := setupEchoApp(deps, true) // true for include GetByID and DeleteByID

	tc := &steps.TodoDeleteByIDContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_delete_by_id.feature"}, tc.InitializeScenario)
}

func TestTodoMarkPendingBDD(t *testing.T) {
	db := setupDatabase(t)
	deps := setupDependencies(db)
	app := setupEchoApp(deps, true) // true for include all routes

	tc := &steps.TodoMarkPendingContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_mark_pending.feature"}, tc.InitializeScenario)
}
