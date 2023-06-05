package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creationDate"`
}

var users []user
var lastId int

func saveUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req user
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
		}
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token inválido",
			})
			return
		}
		lastId++
		req.Id = lastId
		users = append(users, req)
		c.JSON(http.StatusCreated, token)
	}
}

func getUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Users Created": users,
	})
}

func main() {
	router := gin.Default()

	router.POST("/user", func(ctx *gin.Context) {
		var req user
		erro := ctx.ShouldBindJSON(&req)
		if erro != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": erro.Error(),
			})
			return
		}
		req.Id = 1
		ctx.JSON(http.StatusOK, req)
	})
	router.GET("/users", getUsers)
	u := router.Group("/users")
	u.POST("/saved", saveUser())

	router.Run()
}
