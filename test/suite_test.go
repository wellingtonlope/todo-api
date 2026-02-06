package test

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/bootstrap"
	"github.com/wellingtonlope/todo-api/test/steps"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type TestDependencies struct {
	DB *gorm.DB
}

// TestFactory handles test setup using FX bootstrap
type TestFactory struct {
	t       *testing.T
	app     *fx.App
	deps    *TestDependencies
	echoApp *echo.Echo
}

// NewTestFactory creates a new test factory
func NewTestFactory(t *testing.T) *TestFactory {
	return &TestFactory{t: t}
}

// SetupBDDTest returns configured dependencies using FX
func (tf *TestFactory) SetupBDDTest() (*TestDependencies, *echo.Echo) {
	deps := &TestDependencies{}

	// Create FX app with test configuration
	tf.app = fx.New(
		bootstrap.TestFXOptions(),
		fx.Populate(&deps.DB),
		fx.Populate(&tf.echoApp),
	)

	// Start the FX app
	if err := tf.app.Start(context.Background()); err != nil {
		tf.t.Fatalf("Failed to start FX app: %v", err)
	}

	tf.deps = deps
	return deps, tf.echoApp
}

// Reset clears the database between tests
func (tf *TestFactory) Reset(deps *TestDependencies) {
	_ = deps.Reset()
}

// Reset clears all data from the database
func (td *TestDependencies) Reset() error {
	// Simply clear the database - FX maintains all dependencies
	if err := td.DB.Exec("DELETE FROM todos").Error; err != nil {
		return err
	}
	return nil
}

// Cleanup stops the FX app
func (tf *TestFactory) Cleanup() {
	if tf.app != nil {
		_ = tf.app.Stop(context.Background())
	}
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

func TestTodoListBDD(t *testing.T) {
	factory := NewTestFactory(t)
	deps, app := factory.SetupBDDTest()

	tc := &steps.TodoListContext{
		BaseTestContext: steps.BaseTestContext{
			EchoApp: app,
			DB:      deps.DB,
		},
	}

	// Set up the ResetStoreFunc with proper closure
	tc.ResetStoreFunc = func() {
		factory.Reset(deps)
		// Echo app is already configured by FX, no need to recreate
		tc.EchoApp = factory.echoApp
	}

	runBDDTest(t, app, deps.DB, []string{"features/todo_list.feature"}, tc.InitializeScenario)
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
