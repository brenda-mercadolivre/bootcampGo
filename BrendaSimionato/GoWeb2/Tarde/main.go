package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	repository := NewRepository()
	service := NewService(repository)
	user := NewUser(service)

	router := gin.Default()
	users := router.Group("/users")
	users.POST("/post", user.Create())
	users.GET("/get", user.GetAll())

	router.Run()
}
