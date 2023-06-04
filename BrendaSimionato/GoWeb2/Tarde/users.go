package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creationDate"`
}

type Users struct {
	service Service
}

func NewUser(u Service) *Users {
	return &Users{
		service: u,
	}
}

func (c *Users) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token inválido",
			})
			return
		}

		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *Users) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return
		}
		u, err := c.service.Create(req.Name, req.Surname, req.Email, req.Age, req.Height, req.Active, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, u)
	}
}
