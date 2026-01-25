package bootstrap

import (
	"go.uber.org/fx"
)

// FXOptions returns the complete FX configuration for the application
func FXOptions() fx.Option {
	config := Config{
		DatabasePath:  "todo.db",
		WithLifecycle: true,
		WithSwagger:   true,
		Port:          "8080",
	}

	providers := []interface{}{
		// Configuration
		func() Config { return config },

		// Common providers
		provideMiddlewares,
		provideEchoWithLifecycle,
		provideDatabase,
		provideClock,
		provideRepository,
		provideUseCases,
		provideHandlers,
	}

	invokes := []interface{}{
		provideHandlerRegistration(),
		provideSwaggerRegistration(),
	}

	return fx.Options(
		fx.Provide(providers...),
		fx.Invoke(invokes...),
	)
}
