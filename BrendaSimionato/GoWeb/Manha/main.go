package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id          int
	Nome        string
	Sobrenome   string
	Email       string
	Idade       int
	Altura      float64
	Ativo       bool
	DataCriacao string
}

func main() {

	u := user{
		Id:          1,
		Nome:        "Brenda Evellyn",
		Sobrenome:   "Tiba Simionato",
		Email:       "brenda.simionato@mercadolivre.com",
		Idade:       25,
		Altura:      1.58,
		Ativo:       true,
		DataCriacao: "01/06/2023",
	}

	userJSON, erro := json.Marshal(u)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(string(userJSON))

	r := gin.Default()
	r.GET("/name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Hello, Brenda.",
		})
	})
	r.Run()
}
