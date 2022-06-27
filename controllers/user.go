package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/xshot9011/tracing-lab/handlers"
	"github.com/xshot9011/tracing-lab/models"

	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type CreateUserInput struct {
	Name string `form:"name" json:"name"`
	Fibo int    `form:"fibo" json:"fibo`
}

func doFibonacci() uint64 {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	counter := random.Intn(50)

	if counter <= 1 {
		return uint64(counter)
	}

	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < uint(counter); i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}

func AddUser(c *gin.Context) {
	_, span := handlers.Tracer.Start(c.Request.Context(), "AddUser", oteltrace.WithAttributes(attribute.String("id", "1")))
	defer span.End()

	var input CreateUserInput
	input.Name = "Big"
	input.Fibo = int(doFibonacci())

	user := models.User{Name: input.Name, Fibo: input.Fibo}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"status": "200", "Fibo": input.Fibo})
}

func ListUser(c *gin.Context) {
	// _, span := handlers.Tracer.Start(c.Request.Context(), "AddUser", oteltrace.WithAttributes(attribute.String("id", "1")))
	// defer span.End()

	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"status": "200", "users": users})
}
