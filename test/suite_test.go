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
	DB                *gorm.DB
	Clock             usecase.Clock
	Store             interface{}
	CreateUsecase     todo.Create
	GetByIDUsecase    todo.GetByID
	DeleteByIDUsecase todo.DeleteByID
	UpdateUsecase     todo.Update
	CreateHandler     *handler.TodoCreate
	GetByIDHandler    *handler.TodoGetByID
	DeleteByIDHandler *handler.TodoDeleteByID
	UpdateHandler     *handler.TodoUpdate
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
	updateUsecase := todo.NewUpdate(store, clock)
	createHandler := handler.NewTodoCreate(createUsecase)
	getByIDHandler := handler.NewTodoGetByID(getByIDUsecase)
	deleteByIDHandler := handler.NewTodoDeleteByID(deleteByIDUsecase)
	updateHandler := handler.NewTodoUpdate(updateUsecase)

	return &TestDependencies{
		DB:                db,
		Clock:             clock,
		Store:             store,
		CreateUsecase:     createUsecase,
		GetByIDUsecase:    getByIDUsecase,
		DeleteByIDUsecase: deleteByIDUsecase,
		UpdateUsecase:     updateUsecase,
		CreateHandler:     createHandler,
		GetByIDHandler:    getByIDHandler,
		DeleteByIDHandler: deleteByIDHandler,
		UpdateHandler:     updateHandler,
	}
}

func setupEchoApp(deps *TestDependencies, includeGetByID bool, includeUpdate bool) *echo.Echo {
	e := echo.New()
	e.Use(handler.Error)
	e.POST("/todos", deps.CreateHandler.Handle)
	if includeGetByID {
		e.GET("/todos/:id", deps.GetByIDHandler.Handle)
	}
	e.DELETE("/todos/:id", deps.DeleteByIDHandler.Handle)
	if includeUpdate {
		e.PUT("/todos/:id", deps.UpdateHandler.Handle)
	}
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
	app := setupEchoApp(deps, false, false) // false for no GetByID, false for no Update

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
	app := setupEchoApp(deps, true, false) // true for include GetByID, false for no Update

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
	app := setupEchoApp(deps, true, false) // true for include GetByID, false for no Update

	tc := &steps.TodoDeleteByIDContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_delete_by_id.feature"}, tc.InitializeScenario)
}

func TestTodoUpdateBDD(t *testing.T) {
	db := setupDatabase(t)
	deps := setupDependencies(db)
	app := setupEchoApp(deps, true, true) // true for include GetByID, true for Update

	tc := &steps.TodoUpdateContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_update.feature"}, tc.InitializeScenario)
}
