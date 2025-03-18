package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/ppcamp/go-pismo-code-challenge/internal/config"
	"github.com/ppcamp/go-pismo-code-challenge/internal/http/handlers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Serve(ctx context.Context, h *handlers.Handler) error {
	addr := fmt.Sprintf("%s:%d",
		viper.GetString(config.AppHost),
		viper.GetInt(config.AppPort))

	svr := http.Server{
		Addr:        addr,
		Handler:     Routes(h),
		BaseContext: func(listener net.Listener) context.Context { return ctx },
	}

	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()

		// graceful shutdown
		d := viper.GetDuration(config.AppShutdownTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()

		errCh <- svr.Shutdown(ctx)
	}()

	logrus.WithContext(ctx).Infof("Initializing server at http://%s", svr.Addr)
	err := svr.ListenAndServe()
	switch err {
	case http.ErrServerClosed:
		return <-errCh // we need to wait for the shutdown to complete

	default:
		return err
	}
}
