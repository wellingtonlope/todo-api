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
	GetAllUsecase      todo.GetAll
	DeleteByIDUsecase  todo.DeleteByID
	UpdateUsecase      todo.Update
	CompleteUsecase    todo.Complete
	MarkPendingUsecase todo.MarkAsPending

	CreateHandler      *handler.TodoCreate
	GetByIDHandler     *handler.TodoGetByID
	GetAllHandler      *handler.TodoGetAll
	DeleteByIDHandler  *handler.TodoDeleteByID
	UpdateHandler      *handler.TodoUpdate
	CompleteHandler    *handler.TodoComplete
	MarkPendingHandler *handler.TodoMarkPending
}

// TestFactory handles test setup and reduces duplication
type TestFactory struct {
	t *testing.T
}

// NewTestFactory creates a new test factory
func NewTestFactory(t *testing.T) *TestFactory {
	return &TestFactory{t: t}
}

// SetupBDDTest returns configured dependencies
func (tf *TestFactory) SetupBDDTest() (*TestDependencies, *echo.Echo) {
	db := tf.setupDatabase()
	deps := tf.setupDependencies(db)
	app := tf.setupEchoApp(deps)
	return deps, app
}

// Reset clears the state between tests
func (tf *TestFactory) Reset(deps *TestDependencies) {
	_ = deps.Reset()
}

// Reset clears all state including database and store
func (td *TestDependencies) Reset() error {
	// Clear database first
	if err := td.DB.Exec("DELETE FROM todos").Error; err != nil {
		return err
	}
	// Create a fresh store to clear all data
	store := memory.NewTodoRepository()
	td.Store = store
	td.CreateUsecase = todo.NewCreate(store, td.Clock)
	td.GetByIDUsecase = todo.NewGetByID(store)
	td.GetAllUsecase = todo.NewGetAll(store)
	td.DeleteByIDUsecase = todo.NewDeleteByID(store)
	td.UpdateUsecase = todo.NewUpdate(store, td.Clock)
	td.CompleteUsecase = todo.NewComplete(store, td.Clock)
	td.MarkPendingUsecase = todo.NewMarkAsPending(store, td.Clock)
	td.CreateHandler = handler.NewTodoCreate(td.CreateUsecase)
	td.GetByIDHandler = handler.NewTodoGetByID(td.GetByIDUsecase)
	td.GetAllHandler = handler.NewTodoGetAll(td.GetAllUsecase)
	td.DeleteByIDHandler = handler.NewTodoDeleteByID(td.DeleteByIDUsecase)
	td.UpdateHandler = handler.NewTodoUpdate(td.UpdateUsecase)
	td.CompleteHandler = handler.NewTodoComplete(td.CompleteUsecase)
	td.MarkPendingHandler = handler.NewTodoMarkPending(td.MarkPendingUsecase)
	return nil
}

// setupDatabase creates and migrates an in-memory database
func (tf *TestFactory) setupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		tf.t.Fatal(err)
	}
	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		tf.t.Fatal(err)
	}
	return db
}

// setupDependencies creates all use cases and handlers
func (tf *TestFactory) setupDependencies(db *gorm.DB) *TestDependencies {
	clock := clock.NewClientUTC()
	store := memory.NewTodoRepository()
	createUsecase := todo.NewCreate(store, clock)
	getByIDUsecase := todo.NewGetByID(store)
	getAllUsecase := todo.NewGetAll(store)
	deleteByIDUsecase := todo.NewDeleteByID(store)
	updateUsecase := todo.NewUpdate(store, clock)
	completeUsecase := todo.NewComplete(store, clock)
	markPendingUsecase := todo.NewMarkAsPending(store, clock)
	createHandler := handler.NewTodoCreate(createUsecase)
	getByIDHandler := handler.NewTodoGetByID(getByIDUsecase)
	getAllHandler := handler.NewTodoGetAll(getAllUsecase)
	deleteByIDHandler := handler.NewTodoDeleteByID(deleteByIDUsecase)
	updateHandler := handler.NewTodoUpdate(updateUsecase)
	completeHandler := handler.NewTodoComplete(completeUsecase)
	markPendingHandler := handler.NewTodoMarkPending(markPendingUsecase)

	return &TestDependencies{
		DB:                 db,
		Clock:              clock,
		Store:              store,
		CreateUsecase:      createUsecase,
		GetByIDUsecase:     getByIDUsecase,
		GetAllUsecase:      getAllUsecase,
		DeleteByIDUsecase:  deleteByIDUsecase,
		UpdateUsecase:      updateUsecase,
		CompleteUsecase:    completeUsecase,
		MarkPendingUsecase: markPendingUsecase,
		CreateHandler:      createHandler,
		GetByIDHandler:     getByIDHandler,
		GetAllHandler:      getAllHandler,
		DeleteByIDHandler:  deleteByIDHandler,
		UpdateHandler:      updateHandler,
		CompleteHandler:    completeHandler,
		MarkPendingHandler: markPendingHandler,
	}
}

// setupEchoApp configures the Echo app with all handlers
func (tf *TestFactory) setupEchoApp(deps *TestDependencies) *echo.Echo {
	e := echo.New()
	e.Use(handler.Error)
	e.POST("/todos", deps.CreateHandler.Handle)
	e.GET("/todos", deps.GetAllHandler.Handle)
	e.GET("/todos/:id", deps.GetByIDHandler.Handle)
	e.DELETE("/todos/:id", deps.DeleteByIDHandler.Handle)
	e.PUT("/todos/:id", deps.UpdateHandler.Handle)
	e.POST("/todos/:id/complete", deps.CompleteHandler.Handle)
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
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest()

	tc := &steps.TodoCreationContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_creation.feature"}, tc.InitializeScenario)
}

func TestTodoGetByIDBDD(t *testing.T) {
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest()

	tc := &steps.TodoGetByIDContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_get_by_id.feature"}, tc.InitializeScenario)
}

func TestTodoDeleteByIDBDD(t *testing.T) {
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest()

	tc := &steps.TodoDeleteByIDContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_delete_by_id.feature"}, tc.InitializeScenario)
}

func TestTodoGetAllBDD(t *testing.T) {
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest()

	tc := &steps.TodoGetAllContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	// Set up the ResetStoreFunc with proper closure
	tc.ResetStoreFunc = func() {
		factory.Reset(deps)
		newApp := factory.setupEchoApp(deps)
		tc.EchoApp = newApp
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_get_all.feature"}, tc.InitializeScenario)
}

func TestTodoUpdateBDD(t *testing.T) {
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest()

	tc := &steps.TodoUpdateContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_update.feature"}, tc.InitializeScenario)
}

func TestTodoCompleteBDD(t *testing.T) {
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest()

	tc := &steps.TodoCompleteContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_complete.feature"}, tc.InitializeScenario)
}

func TestTodoMarkPendingBDD(t *testing.T) {
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest()

	tc := &steps.TodoMarkPendingContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_mark_pending.feature"}, tc.InitializeScenario)
}
