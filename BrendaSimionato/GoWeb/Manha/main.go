package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id           int
	Name         string
	Surname      string
	Email        string
	Age          int
	Height       float64
	Active       bool
	CreationDate string
}

func UsersFromJSON(file string) (users []user, err error) {
	data, err := os.ReadFile("GoWeb/users.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return
	}
	return
}

func main() {

	router := gin.Default()
	router.GET("/name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Hello, Brenda.",
		})
	})

	router.GET("/users", func(c *gin.Context) {
		users, err := UsersFromJSON("GoWeb/users.json")
		if err != nil {
			c.Status(500)
			return
		}
		c.JSON(200, gin.H{"data": users})
	})

	router.Run()
}
