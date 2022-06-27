package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/xshot9011/tracing-lab/models"
)

type CreateUserInput struct {
	Name string `form:"name"`
	Fibo int    `form:"fibo"`
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
	var input CreateUserInput
	input.Name = "Big"
	input.Fibo = int(doFibonacci())

	user := models.User{Name: input.Name, Fibo: input.Fibo}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"Status": "200", "Fibo": input.Fibo})
}
