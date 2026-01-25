package bootstrap

// Config holds environment-specific configuration for bootstrap
type Config struct {
	DatabasePath  string // Path to SQLite database file
	WithLifecycle bool   // Whether to add lifecycle hooks to Echo
	WithSwagger   bool   // Whether to add Swagger documentation
	Port          string // Port for Echo server (used only with lifecycle)
}
