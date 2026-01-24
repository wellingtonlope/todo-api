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

// FeatureType represents the type of BDD test feature
type FeatureType int

const (
	FeatureCreation FeatureType = iota
	FeatureGetByID
	FeatureGetAll
	FeatureDeleteByID
	FeatureUpdate
	FeatureComplete
	FeatureMarkPending
)

// FeatureFlags defines which endpoints to include in the Echo app
type FeatureFlags struct {
	IncludeGetByID     bool
	IncludeGetAll      bool
	IncludeUpdate      bool
	IncludeComplete    bool
	IncludeMarkPending bool
}

// GetFeatureFlags returns the appropriate feature flags for a given feature type
func GetFeatureFlags(featureType FeatureType) FeatureFlags {
	switch featureType {
	case FeatureCreation:
		return FeatureFlags{}
	case FeatureGetByID:
		return FeatureFlags{IncludeGetByID: true}
	case FeatureGetAll:
		return FeatureFlags{IncludeGetAll: true}
	case FeatureDeleteByID:
		return FeatureFlags{IncludeGetByID: true}
	case FeatureUpdate:
		return FeatureFlags{IncludeGetByID: true, IncludeUpdate: true}
	case FeatureComplete:
		return FeatureFlags{IncludeGetByID: true, IncludeComplete: true}
	case FeatureMarkPending:
		return FeatureFlags{IncludeGetByID: true, IncludeMarkPending: true}
	default:
		return FeatureFlags{}
	}
}

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

// SetupBDDTest returns configured dependencies for a given feature type
func (tf *TestFactory) SetupBDDTest(featureType FeatureType) (*TestDependencies, *echo.Echo) {
	db := tf.setupDatabase()
	deps := tf.setupDependencies(db)
	app := tf.setupEchoApp(deps, GetFeatureFlags(featureType))
	return deps, app
}

// Reset clears the state between tests
func (tf *TestFactory) Reset(deps *TestDependencies) {
	deps.ResetStore()
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
	td.CompleteUsecase = todo.NewComplete(store, td.Clock)
	td.MarkPendingUsecase = todo.NewMarkAsPending(store, td.Clock)
	td.CreateHandler = handler.NewTodoCreate(td.CreateUsecase)
	td.GetByIDHandler = handler.NewTodoGetByID(td.GetByIDUsecase)
	td.GetAllHandler = handler.NewTodoGetAll(td.GetAllUsecase)
	td.DeleteByIDHandler = handler.NewTodoDeleteByID(td.DeleteByIDUsecase)
	td.UpdateHandler = handler.NewTodoUpdate(td.UpdateUsecase)
	td.CompleteHandler = handler.NewTodoComplete(td.CompleteUsecase)
	td.MarkPendingHandler = handler.NewTodoMarkPending(td.MarkPendingUsecase)
}

// Reset clears all state including database and store
func (td *TestDependencies) Reset() error {
	// Clear database
	if err := td.DB.Exec("DELETE FROM todos").Error; err != nil {
		return err
	}
	// Reset store
	td.ResetStore()
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

// setupEchoApp configures the Echo app with the given feature flags
func (tf *TestFactory) setupEchoApp(deps *TestDependencies, flags FeatureFlags) *echo.Echo {
	e := echo.New()
	e.Use(handler.Error)
	e.POST("/todos", deps.CreateHandler.Handle)
	if flags.IncludeGetAll {
		e.GET("/todos", deps.GetAllHandler.Handle)
	}
	if flags.IncludeGetByID {
		e.GET("/todos/:id", deps.GetByIDHandler.Handle)
	}
	e.DELETE("/todos/:id", deps.DeleteByIDHandler.Handle)
	if flags.IncludeUpdate {
		e.PUT("/todos/:id", deps.UpdateHandler.Handle)
	}
	if flags.IncludeComplete {
		e.POST("/todos/:id/complete", deps.CompleteHandler.Handle)
	}
	if flags.IncludeMarkPending {
		e.POST("/todos/:id/pending", deps.MarkPendingHandler.Handle)
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
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest(FeatureCreation)

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
	deps, app := factory.SetupBDDTest(FeatureGetByID)

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
	deps, app := factory.SetupBDDTest(FeatureDeleteByID)

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
	deps, app := factory.SetupBDDTest(FeatureGetAll)

	tc := &steps.TodoGetAllContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	// Set up the ResetStoreFunc with proper closure
	tc.ResetStoreFunc = func() {
		factory.Reset(deps)
		newApp := factory.setupEchoApp(deps, GetFeatureFlags(FeatureGetAll))
		tc.EchoApp = newApp
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_get_all.feature"}, tc.InitializeScenario)
}

func TestTodoUpdateBDD(t *testing.T) {
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest(FeatureUpdate)

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
	deps, app := factory.SetupBDDTest(FeatureComplete)

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
	deps, app := factory.SetupBDDTest(FeatureMarkPending)

	tc := &steps.TodoMarkPendingContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_mark_pending.feature"}, tc.InitializeScenario)
}
