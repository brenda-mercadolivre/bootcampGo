package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

var USERS_FILE = "GoWeb/users.json"

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

func Filter[T any](arr []T, predicate func(T) bool) (result []T) {
	for _, x := range arr {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return
}

func MatchUserWithParameters(c *gin.Context) func(user) bool {
	return nil
}

func SliceId(c *gin.Context) {
	users, err := UsersFromJSON(USERS_FILE)
	if err != nil {
		c.Status(500)
		return
	}
	var filtered []user
	for _, user := range users {
		filtered = append(filtered, user)
	}
	c.JSON(200, filtered)
}

func SliceName(c *gin.Context) {
	users, err := UsersFromJSON(USERS_FILE)
	if err != nil {
		c.Status(500)
		return
	}
	filtered := []string{}
	for _, user := range users {
		filtered = append(filtered, user.Name)
	}
	c.JSON(200, filtered)
}
func SliceEmail(c *gin.Context) {
	users, err := UsersFromJSON(USERS_FILE)
	if err != nil {
		c.Status(500)
		return
	}
	filtered := []string{}
	for _, user := range users {
		filtered = append(filtered, user.Email)
	}
	c.JSON(200, filtered)
}

func main() {
	router := gin.Default()
	users := router.Group("/users")
	{
		users.GET("/id", SliceId)
		users.GET("/name", SliceName)
		users.GET("/email", SliceEmail)
	}
	router.Run()
}
