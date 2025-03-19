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

	db, err := db.New(ctx, db.Params{
		Host:     viper.GetString(config.DatabaseHost),
		Port:     viper.GetInt(config.DatabasePort),
		Driver:   viper.GetString(config.DatabaseDriver),
		User:     viper.GetString(config.DatabaseUsername),
		Password: viper.GetString(config.DatabasePassword),
	})
	if err != nil {
		logrus.WithError(err).Fatal("fail to connect to db: %w", err)
	}

	h := &handlers.Handler{
		Account: services.NewAccountService(db),
	}

	err = http.Serve(ctx, h)
	if err == nil {
		logrus.Info("Server shutdown successfully")
	} else {
		logrus.WithError(err).Fatal("Fail to serve http server")
	}
}
