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
	GetAllUsecase     todo.GetAll
	DeleteByIDUsecase todo.DeleteByID
	UpdateUsecase     todo.Update
	CreateHandler     *handler.TodoCreate
	GetByIDHandler    *handler.TodoGetByID
	GetAllHandler     *handler.TodoGetAll
	DeleteByIDHandler *handler.TodoDeleteByID
	UpdateHandler     *handler.TodoUpdate
}

func (td *TestDependencies) ResetStore() {
	// Create a new store to clear all data
	store := memory.NewTodoRepository()
	td.Store = store
	td.CreateUsecase = todo.NewCreate(store, td.Clock)
	td.GetByIDUsecase = todo.NewGetByID(store)
	td.GetAllUsecase = todo.NewGetAll(store)
	td.DeleteByIDUsecase = todo.NewDeleteByID(store)
	td.UpdateUsecase = todo.NewUpdate(store, td.Clock)
	td.CreateHandler = handler.NewTodoCreate(td.CreateUsecase)
	td.GetByIDHandler = handler.NewTodoGetByID(td.GetByIDUsecase)
	td.GetAllHandler = handler.NewTodoGetAll(td.GetAllUsecase)
	td.DeleteByIDHandler = handler.NewTodoDeleteByID(td.DeleteByIDUsecase)
	td.UpdateHandler = handler.NewTodoUpdate(td.UpdateUsecase)
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
	getAllUsecase := todo.NewGetAll(store)
	deleteByIDUsecase := todo.NewDeleteByID(store)
	updateUsecase := todo.NewUpdate(store, clock)
	createHandler := handler.NewTodoCreate(createUsecase)
	getByIDHandler := handler.NewTodoGetByID(getByIDUsecase)
	getAllHandler := handler.NewTodoGetAll(getAllUsecase)
	deleteByIDHandler := handler.NewTodoDeleteByID(deleteByIDUsecase)
	updateHandler := handler.NewTodoUpdate(updateUsecase)

	return &TestDependencies{
		DB:                db,
		Clock:             clock,
		Store:             store,
		CreateUsecase:     createUsecase,
		GetByIDUsecase:    getByIDUsecase,
		GetAllUsecase:     getAllUsecase,
		DeleteByIDUsecase: deleteByIDUsecase,
		UpdateUsecase:     updateUsecase,
		CreateHandler:     createHandler,
		GetByIDHandler:    getByIDHandler,
		GetAllHandler:     getAllHandler,
		DeleteByIDHandler: deleteByIDHandler,
		UpdateHandler:     updateHandler,
	}
}

func setupEchoApp(deps *TestDependencies, includeGetByID, includeGetAll, includeUpdate, includeMarkPending bool) *echo.Echo {
	e := echo.New()
	e.Use(handler.Error)
	e.POST("/todos", deps.CreateHandler.Handle)
	if includeGetAll {
		e.GET("/todos", deps.GetAllHandler.Handle)
	}
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
	app := setupEchoApp(deps, false, false, false, false) // false for no GetByID, false for no GetAll, false for no Update

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
	app := setupEchoApp(deps, true, false, false, false) // true for include GetByID, false for no GetAll, false for no Update

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
	app := setupEchoApp(deps, true, false, false, false) // true for include GetByID, false for no GetAll, false for no Update

	tc := &steps.TodoDeleteByIDContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_delete_by_id.feature"}, tc.InitializeScenario)
}

func TestTodoGetAllBDD(t *testing.T) {
	db := setupDatabase(t)
	deps := setupDependencies(db)
	app := setupEchoApp(deps, false, true, false, false) // false for no GetByID, true for include GetAll, false for no Update

	tc := &steps.TodoGetAllContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
		ResetStoreFunc: func() {
			deps.ResetStore()
		},
	}

	// Update the ResetStoreFunc to also update the EchoApp after tc is created
	tc.ResetStoreFunc = func() {
		deps.ResetStore()
		newApp := setupEchoApp(deps, false, true, false, false)
		tc.EchoApp = newApp
	}

	runBDDTest(t, app, db, []string{"features/todo_get_all.feature"}, tc.InitializeScenario)
}

func TestTodoUpdateBDD(t *testing.T) {
	db := setupDatabase(t)
	deps := setupDependencies(db)
	app := setupEchoApp(deps, true, false, true, false) // true for include GetByID, false for no GetAll, true for Update

	tc := &steps.TodoUpdateContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_update.feature"}, tc.InitializeScenario)
}

func TestTodoMarkPendingBDD(t *testing.T) {
	db := setupDatabase(t)
	deps := setupDependencies(db)
	app := setupEchoApp(deps, true, false, false, true) // true for include GetByID, false for no GetAll, false for no Update, true for MarkPending

	tc := &steps.TodoMarkPendingContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      db,
		},
	}

	runBDDTest(t, app, db, []string{"features/todo_mark_pending.feature"}, tc.InitializeScenario)
}
