package config

// App related viper identifiers
const (
	AppName               = "app.name"
	AppPort               = "app.port"
	AppHost               = "app.host"
	AppShutdownTimeout    = "app.timeout"
	AppCorsAllowedOrigins = "app.cors.allowed_origins"
	AppCorsAllowedMethods = "app.cors.allowed_methods"
)

// Logging related viper identifiers
const (
	LoggingLevel       = "logging.level"
	LoggingHttpEnabled = "logging.level"
)

// Database related viper identifiers
const (
	DatabaseDriver   = "database.driver"
	DatabaseHost     = "database.host"
	DatabasePort     = "database.port"
	DatabaseUsername = "database.username"
	DatabasePassword = "database.password"
)
