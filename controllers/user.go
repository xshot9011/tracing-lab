package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/xshot9011/tracing-lab/models"
)

type CreateUserInput struct {
	Name string `form:"name"`
	Fibo int    `form:"fibo"`
}

func AddUser(c *gin.Context) {
	var input CreateUserInput
	input.Name = c.PostForm("name")
	input.Fibo, _ = strconv.Atoi(c.PostForm("fibo"))

	user := models.Book{Name: input.Name, Fibo: input.Fibo}
	models.User.Create(&user)

	c.JSON(http.StatusOK, gin.H{"Status": "200"})
}
