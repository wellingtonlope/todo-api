package bootstrap

import (
	"os"

	"go.uber.org/fx"
)

// FXOptions returns the complete FX configuration for the application
func FXOptions() fx.Option {
	return fx.Options(
		// Production configuration
		fx.Supply(Config{
			Database: DatabaseConfig{
				Driver:   getEnv("DB_DRIVER", "mysql"),
				Host:     getEnv("DB_HOST", "localhost"),
				Port:     getEnv("DB_PORT", "3306"),
				User:     getEnv("DB_USER", "todo_user"),
				Password: getEnv("DB_PASSWORD", "todo_password"),
				Database: getEnv("DB_NAME", "todo_api"),
			},
			WithLifecycle: true,
			WithSwagger:   true,
			Port:          getEnv("PORT", "8080"),
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

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
