package telemetry

import (
    "context"
    "errors"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/metric"
    "go.opentelemetry.io/otel/sdk/trace"
    "github.com/jupitercloud/subscribed/logger"
)

var log = logger.Named("telemetry");

type ExportMode int;

const (
	ExportModeNone = iota
    ExportModeConsole
    ExportModeGrpc
)

func ExportModeFromString(mode string) ExportMode {
    if mode == "none" { return ExportModeNone }
    if mode == "console" { return ExportModeConsole }
    if mode == "grpc" { return ExportModeGrpc }
    log.Warn("Invalid telemetry export mode", "mode", mode)
    return ExportModeNone
}

// Initialize bootstraps the OpenTelemetry pipeline and returns a shutdown callback.
func Initialize(ctx context.Context, exportMode ExportMode) (shutdown func(context.Context) error, err error) {
    if exportMode == ExportModeNone {
        return func(ctx context.Context) error { return nil }, nil
    }

    log.Info("Initializing OpenTelemetry")
    var shutdownFuncs []func(context.Context) error

    // shutdown calls cleanup functions registered via shutdownFuncs.
    // The errors from the calls are joined.
    // Each registered cleanup will be invoked once.
    shutdown = func(ctx context.Context) error {
        log.Info("Shutting down OpenTelemetry")
        var err error
        for _, fn := range shutdownFuncs {
            err = errors.Join(err, fn(ctx))
        }
        shutdownFuncs = nil
        return err
    }

    // handleErr calls shutdown for cleanup and makes sure that all errors are returned.
    handleErr := func(inErr error) {
        err = errors.Join(inErr, shutdown(ctx))
    }

    // Set up propagator.
    prop := newPropagator()
    otel.SetTextMapPropagator(prop)

    // Set up trace provider.
    tracerProvider, err := newTraceProvider(exportMode)
    if err != nil {
        handleErr(err)
        return nil, err
    }

    shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
    otel.SetTracerProvider(tracerProvider)

    // Set up meter provider.
    meterProvider, err := newMeterProvider(exportMode)
    if err != nil {
        handleErr(err)
        return nil, err
    }
    shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
    otel.SetMeterProvider(meterProvider)

    return shutdown, nil
}

func newPropagator() propagation.TextMapPropagator {
    return propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{},
        propagation.Baggage{},
    )
}

func newTraceProvider(exportMode ExportMode) (*trace.TracerProvider, error) {
    var err error = nil
    var exporter trace.SpanExporter = nil
    switch (exportMode) {
    case ExportModeConsole:
        exporter, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
        break;
    case ExportModeGrpc:
        exporter, err = otlptracegrpc.New(context.Background())
        break;
    }

    if err != nil {
        return nil, err
    }

    traceProvider := trace.NewTracerProvider(trace.WithBatcher(exporter))
    return traceProvider, nil
}

func newMeterProvider(exportMode ExportMode) (*metric.MeterProvider, error) {
    var err error = nil
    var exporter metric.Exporter = nil
    switch (exportMode) {
    case ExportModeConsole:
        exporter, err = stdoutmetric.New()
        break;
    case ExportModeGrpc:
        exporter, err = otlpmetricgrpc.New(context.Background())
        break;
    }

    if err != nil {
        return nil, err
    }

    meterProvider := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(exporter)))
    return meterProvider, nil
}
