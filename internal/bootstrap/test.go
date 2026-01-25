package bootstrap

import (
	"go.uber.org/fx"
)

// TestFXOptions returns FX configuration for BDD tests using in-memory SQLite
func TestFXOptions() fx.Option {
	config := Config{
		DatabasePath:  ":memory:",
		WithLifecycle: false,
		WithSwagger:   false,
		Port:          "",
	}

	providers := []interface{}{
		// Configuration
		func() Config { return config },

		// Common providers
		provideMiddlewares,
		provideEcho,
		provideDatabase,
		provideClock,
		provideRepository,
		provideUseCases,
		provideHandlers,
	}

	invokes := []interface{}{
		provideHandlerRegistration(),
	}

	return fx.Options(
		fx.Provide(providers...),
		fx.Invoke(invokes...),
	)
}
