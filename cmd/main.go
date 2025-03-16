package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ppcamp/go-pismo-code-challenge/internal/config"
	"github.com/ppcamp/go-pismo-code-challenge/internal/http"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/utils/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := config.LoadViperConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Fail to load viper configs")
	}

	err = logging.SetupLogrus(viper.GetString(config.LoggingLevel))
	if err != nil {
		logrus.WithError(err).Fatal("fail to configure logging: %w", err)
	}

	err = http.Serve(ctx)
	if err == nil {
		logrus.Info("Server shutdown successfully")
	} else {
		logrus.WithError(err).Fatal("Fail to serve http server")
	}
}
