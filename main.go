package main

import (
	"github.com/xshot9011/tracing-lab/controllers"
	"github.com/xshot9011/tracing-lab/models"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())

	router.POST("/", controllers.AddUser)

	return router
}

func main() {
	router := setupRouter()
	models.InitDatabase()

	router.Run(":80")
}
