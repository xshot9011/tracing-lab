package handlers

import (
	"context"
	"os"

	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

var (
	OTEL_EXPORTER_OTLP_ENDPOINT = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	SERVICE_NAME                = os.Getenv("SERVICE_NAME")
	Tracer                      = otel.Tracer("gin-server")
)

func InitTracer() (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	endpoint := OTEL_EXPORTER_OTLP_ENDPOINT
	if endpoint == "" {
		endpoint = "0.0.0.0:4317"
	}

	// Create and start new OTLP trace exporter (TLS config)
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(endpoint), otlptracegrpc.WithDialOption(grpc.WithBlock()))
	HandleErr(err, "Failed to create new OTLP trace exporter")

	idg := xray.NewIDGenerator()

	service := SERVICE_NAME
	if service == "" {
		service = "go-application"
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		// the service name used to display traces in backends
		semconv.ServiceNameKey.String("test-service"),
	)
	HandleErr(err, "failed to create resource")

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithIDGenerator(idg),
	)

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(xray.Propagator{})

	return traceProvider, nil
}
