package bootstrap

import (
	"go.uber.org/fx"
)

// FXOptions returns the complete FX configuration for the application
func FXOptions() fx.Option {
	return fx.Options(
		// Production configuration
		fx.Supply(Config{
			DatabasePath:  "todo.db",
			WithLifecycle: true,
			WithSwagger:   true,
			Port:          "8080",
		}),
		// Infrastructure providers (middlewares, database, handler registration)
		InfrastructureProviders(),
		// Common providers (clock, repositories, use cases, handlers)
		CommonProviders(),
		// Production-specific providers
		fx.Provide(provideEchoWithLifecycle),
		// Production-specific invokes
		fx.Invoke(provideSwaggerRegistration()),
	)
}
