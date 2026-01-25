package bootstrap

import (
	"go.uber.org/fx"
)

// InfrastructureProviders returns fx.Option with shared infrastructure providers
func InfrastructureProviders() fx.Option {
	providers := []interface{}{
		// Common providers
		provideMiddlewares,
		provideDatabase,
	}

	invokes := []interface{}{
		provideHandlerRegistration(),
	}

	return fx.Module("infrastructure",
		fx.Provide(providers...),
		fx.Invoke(invokes...),
	)
}
