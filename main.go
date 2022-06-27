package main

// https://aws-otel.github.io/docs/getting-started/go-sdk/trace-manual-instr
// https://github.com/aws-observability/aws-otel-go/blob/main/sampleapp/main.go
// https://signoz.io/blog/opentelemetry-gin/

import (
	"net/http"

	"github.com/xshot9011/tracing-lab/controllers"
	"github.com/xshot9011/tracing-lab/handlers"
	"github.com/xshot9011/tracing-lab/models"

	"github.com/gin-gonic/gin"
	middleware "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.Use(middleware.Middleware("application"))

	router.GET("/", controllers.AddUser)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "200",
		})
	})

	router.GET("/users", controllers.ListUser)

	return router
}

func main() {
	handlers.InitLogConfiguration()
	// tp, err := handlers.InitTracer()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err := tp.Shutdown(context.Background()); err != nil {
	// 		log.Printf("Error shutting down tracer provider: %v", err)
	// 	}
	// }()
	models.InitDatabase()

	router := setupRouter()
	router.Run(":80")
}
