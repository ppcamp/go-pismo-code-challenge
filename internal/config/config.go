package config

// App related viper identifiers
const (
	AppName               = "service.name"
	AppPort               = "app.port"
	AppHost               = "app.host"
	AppShutdownTimeout    = "app.timeout"
	AppCorsAllowedOrigins = "app.cors.allowed_origins"
	AppCorsAllowedMethods = "app.cors.allowed_methods"
)

// Logging related viper identifiers
const (
	LoggingLevel       = "logging.level"
	LoggingHttpEnabled = "logging.http_enabled"
)

// Database related viper identifiers
const (
	DatabaseDriver   = "database.driver"
	DatabaseHost     = "database.host"
	DatabasePort     = "database.port"
	DatabaseUsername = "database.username"
	DatabasePassword = "database.password"
)
