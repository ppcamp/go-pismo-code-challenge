package config

// App related viper identifiers
const (
	AppPort            = "app.port"
	AppHost            = "app.host"
	AppShutdownTimeout = "app.timeout"
)

// Logging related viper identifiers
const (
	LoggingLevel       = "logging.level"
	LoggingHttpEnabled = "logging.level"
)

// Database related viper identifiers
const (
	DatabaseHost     = "database.host"
	DatabasePort     = "database.port"
	DatabaseUsername = "database.username"
	DatabasePassword = "database.password"
)
