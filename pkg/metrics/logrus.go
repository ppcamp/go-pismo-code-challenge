package metrics

import (
	"github.com/sirupsen/logrus"

	"go.opentelemetry.io/contrib/bridges/otellogrus"
	"go.opentelemetry.io/otel/log"
)

// Logrus setup otel for logrus globally.
//
// Example
//
//	Logrus(viper.GetString(config.AppName))
func LogrusGlobal(name string, provider log.LoggerProvider) error {
	hook := otellogrus.NewHook(name,
		otellogrus.WithLoggerProvider(provider))

	logrus.AddHook(hook)

	return nil
}

// Logrus setup otel for logrus globally.
//
// Example
//
//	Logrus(viper.GetString(config.AppName))
func Logrus(log *logrus.Logger, name string, provider log.LoggerProvider) error {
	hook := otellogrus.NewHook(name,
		otellogrus.WithLoggerProvider(provider))

	log.AddHook(hook)

	return nil
}
