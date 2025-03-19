package metrics

import (
	"context"
	"fmt"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
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
		if err = m.initLogs(ctx); err != nil {
			return
		}

		if err = m.initTraces(ctx); err != nil {
			return
		}

		if err = m.initMetrics(ctx); err != nil {
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

// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp#example-package
func (m *management) initLogs(ctx context.Context) error {
	exp, err := otlploghttp.New(ctx)
	if err != nil {
		return fmt.Errorf("fail to initialize otel log: %w", err)
	}

	processor := log.NewBatchProcessor(exp)
	provider := log.NewLoggerProvider(log.WithProcessor(processor))
	m.shutdownCbs = append(m.shutdownCbs, provider.Shutdown)

	global.SetLoggerProvider(provider)
	return nil
}

// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp#example-package
func (m *management) initTraces(ctx context.Context) error {
	exp, err := otlptracehttp.New(ctx)
	if err != nil {
		return fmt.Errorf("fail to initialize otel tracers: %w", err)
	}

	tracerProvider := trace.NewTracerProvider(trace.WithBatcher(exp))
	m.shutdownCbs = append(m.shutdownCbs, tracerProvider.Shutdown)

	otel.SetTracerProvider(tracerProvider)
	return nil
}

// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp#example-package
func (m *management) initMetrics(ctx context.Context) error {
	exp, err := otlpmetrichttp.New(ctx)
	if err != nil {
		panic(err)
	}

	meterProvider := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(exp)))
	m.shutdownCbs = append(m.shutdownCbs, meterProvider.Shutdown)

	otel.SetMeterProvider(meterProvider)
	return nil
}
