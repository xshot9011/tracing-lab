package main

// https://aws-otel.github.io/docs/getting-started/go-sdk/trace-manual-instr
// https://github.com/aws-observability/aws-otel-go/blob/main/sampleapp/main.go
// https://signoz.io/blog/opentelemetry-gin/

import (
	"context"
	"os"

	"github.com/xshot9011/tracing-lab/controllers"
	"github.com/xshot9011/tracing-lab/handlers"
	"github.com/xshot9011/tracing-lab/models"

	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	middleware "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var (
	OTEL_EXPORTER_OTLP_ENDPOINT = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	SERVICE_NAME                = os.Getenv("SERVICE_NAME")
	// tracer                      = otel.Tracer("application")
)

func initTracer() {
	ctx := context.Background()

	endpoint := OTEL_EXPORTER_OTLP_ENDPOINT
	if endpoint == "" {
		endpoint = "0.0.0.0:4317"
	}

	// Create and start new OTLP trace exporter (TLS config)
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(endpoint), otlptracegrpc.WithDialOption(grpc.WithBlock()))
	handlers.HandleErr(err, "Failed to create new OTLP trace exporter")

	idg := xray.NewIDGenerator()

	service := SERVICE_NAME
	if service == "" {
		service = "go-gorilla"
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		// the service name used to display traces in backends
		semconv.ServiceNameKey.String("test-service"),
	)
	handlers.HandleErr(err, "failed to create resource")

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithIDGenerator(idg),
	)

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(xray.Propagator{})
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.Use(middleware.Middleware("application"))

	router.GET("/", controllers.AddUser)

	return router
}

func main() {
	handlers.InitLogConfiguration()
	// initProvider()
	models.InitDatabase()

	router := setupRouter()
	router.Run(":80")
}
