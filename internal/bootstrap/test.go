package bootstrap

import (
	"go.uber.org/fx"
)

// TestFXOptions returns FX configuration for BDD tests using in-memory SQLite
func TestFXOptions() fx.Option {
	return fx.Options(
		// Test configuration
		fx.Supply(Config{
			Database: DatabaseConfig{
				Driver: "sqlite",
				Path:   ":memory:",
			},
			WithLifecycle: false,
			WithSwagger:   false,
			Port:          "",
		}),
		// Infrastructure providers (middlewares, database, handler registration)
		InfrastructureProviders(),
		// Common providers (clock, repositories, use cases, handlers)
		CommonProviders(),
		// Test-specific providers
		fx.Provide(provideEcho),
	)
}
