package main

import (
	"fmt"
	"net/http"
	"strconv"

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

func (c *Users) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "token inválido"})
			return
		}

		Id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "ID inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "O nome do usuário é obrigatório"})
			return
		}

		if req.Surname == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "O sobrenome do usuário é obrigatório"})
			return
		}
		if req.Email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "O email do usuário é obrigatório"})
			return
		}
		if req.Age == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "A idade do usuário é obrigatória"})
			return
		}
		if req.Height == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "A altura do usuário é obrigatória"})
			return
		}
		if req.Active == false {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "O campo é ativo, do usuário é obrigatório"})
			return
		}
		if req.CreationDate == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "A data de criação do usuário é obrigatória"})
			return
		}

		u, err := c.service.Update(int(Id), req.Name, req.Surname, req.Email, req.Age, req.Height, req.Active, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *Users) UpdateSurnameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "ID inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		if req.Surname == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "O sobrenome do usuário é obrigatório"})
			return
		}
		if req.Age == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "A idade do usuário é obrigatória"})
			return
		}

		u, err := c.service.UpdateSurnameAndAge(int(id), req.Surname, req.Age)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *Users) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "ID inválido"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"Data": fmt.Sprintf("O usuário %d foi removido", id)})
	}
}
