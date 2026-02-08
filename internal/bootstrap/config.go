package bootstrap

// Config holds environment-specific configuration for bootstrap
type Config struct {
	Database      DatabaseConfig // MySQL database configuration
	WithLifecycle bool           // Whether to add lifecycle hooks to Echo
	WithSwagger   bool           // Whether to add Swagger documentation
	Port          string         // Port for Echo server (used only with lifecycle)
}

// DatabaseConfig holds MySQL connection configuration
type DatabaseConfig struct {
	Driver   string // Database driver (mysql, sqlite)
	Host     string // Database host
	Port     string // Database port
	User     string // Database username
	Password string // Database password
	Database string // Database name
	Path     string // Path to SQLite database file (for tests)
}
