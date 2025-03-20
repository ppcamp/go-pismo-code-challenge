package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ppcamp/go-pismo-code-challenge/internal/config"
	"github.com/ppcamp/go-pismo-code-challenge/internal/http"
	"github.com/ppcamp/go-pismo-code-challenge/internal/http/handlers"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	"github.com/ppcamp/go-pismo-code-challenge/internal/services"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/metrics"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/utils/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/log/global"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := config.LoadViperConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Fail to load viper configs")
	}

	err = metrics.Init(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("fail to initialize metrics")
	}

	err = logging.LogrusGlobal(viper.GetString(config.LoggingLevel))
	if err != nil {
		logrus.WithError(err).Fatal("fail to setup logger: %w", err)
	}

	appName := viper.GetString(config.AppName)
	err = metrics.LogrusGlobal(appName, global.GetLoggerProvider())
	if err != nil {
		logrus.WithError(err).Fatal("fail to configure otel log: %w", err)
	}

	db, err := db.New(ctx, db.Params{
		Host:     viper.GetString(config.DatabaseHost),
		Port:     viper.GetInt(config.DatabasePort),
		Driver:   viper.GetString(config.DatabaseDriver),
		User:     viper.GetString(config.DatabaseUsername),
		Password: viper.GetString(config.DatabasePassword),
		DB:       viper.GetString(config.DatabaseDb),
	})
	if err != nil {
		logrus.WithError(err).Fatal("fail to connect to db: %w", err)
	}
	defer db.Close(ctx)

	h := &handlers.Handler{
		Account:     services.NewAccountService(db),
		Transaction: services.NewTransactionService(db),
	}

	err = http.Serve(ctx, h)
	if err == nil {
		logrus.Info("Server shutdown successfully")
	} else {
		logrus.WithError(err).Fatal("Fail to serve http server")
	}
}
