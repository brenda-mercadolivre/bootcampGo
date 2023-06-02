package main

import (
	"encoding/json"
	"os"
	"strconv"

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

func GetById(ctxt *gin.Context) {
	users, err := UsersFromJSON(USERS_FILE)
	if err != nil {
		ctxt.Status(500)
		return
	}

	u := ctxt.Param("id")
	uToInt, err := strconv.Atoi(u)
	for _, user := range users {
		if user.Id == uToInt {
			ctxt.JSON(200, user)
			return
		}
	}
	ctxt.Status(404)
}

func main() {
	router := gin.Default()
	router.GET("/user/:id", GetById)

	router.Run()
}
