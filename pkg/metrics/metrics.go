package metrics

import (
	"context"
	"fmt"
	"sync"

	"github.com/ppcamp/go-pismo-code-challenge/internal/config"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type shutdownFunc func(context.Context) error

type management struct {
	start, close sync.Once

	shutdownCbs []shutdownFunc
}

var m management

func Init(ctx context.Context) error {
	var err error

	m.start.Do(func() {
		var res *resource.Resource
		res, err = m.resource(ctx)
		if err != nil {
			return
		}

		if err = m.initLogs(ctx, res); err != nil {
			return
		}

		if err = m.initTraces(ctx, res); err != nil {
			return
		}

		if err = m.initMetrics(ctx, res); err != nil {
			return
		}
	})

	return err
}

func Shutdown(ctx context.Context) error {
	errArr := make([]error, 0, len(m.shutdownCbs))

	m.close.Do(func() {
		for _, shutdown := range m.shutdownCbs {
			if err := shutdown(ctx); err != nil {
				errArr = append(errArr, err)
			}
		}
	})

	if len(errArr) > 0 {
		return fmt.Errorf("fail to shutdown: %v", errArr)
	}

	return nil
}

func (m *management) resource(ctx context.Context) (*resource.Resource, error) {
	return resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(viper.GetString(config.AppName)),
		),
	)
}

// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp#example-package
func (m *management) initLogs(ctx context.Context, res *resource.Resource) error {
	exp, err := otlploghttp.New(ctx)
	if err != nil {
		return fmt.Errorf("fail to initialize otel log: %w", err)
	}

	processor := log.NewBatchProcessor(exp)
	provider := log.NewLoggerProvider(
		log.WithProcessor(processor),
		log.WithResource(res))

	m.shutdownCbs = append(m.shutdownCbs, provider.Shutdown)

	global.SetLoggerProvider(provider)
	return nil
}

// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp#example-package
func (m *management) initTraces(ctx context.Context, res *resource.Resource) error {
	exp, err := otlptracehttp.New(ctx)
	if err != nil {
		return fmt.Errorf("fail to initialize otel tracers: %w", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(res))
	m.shutdownCbs = append(m.shutdownCbs, tracerProvider.Shutdown)

	otel.SetTracerProvider(tracerProvider)
	return nil
}

// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp#example-package
func (m *management) initMetrics(ctx context.Context, res *resource.Resource) error {
	exp, err := otlpmetrichttp.New(ctx)
	if err != nil {
		panic(err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exp)),
		metric.WithResource(res))
	m.shutdownCbs = append(m.shutdownCbs, meterProvider.Shutdown)

	otel.SetMeterProvider(meterProvider)
	return nil
}
