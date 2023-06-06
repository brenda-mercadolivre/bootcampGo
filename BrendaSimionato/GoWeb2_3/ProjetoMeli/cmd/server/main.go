package main

import (
	"log"

	"github.com/brenda-mercadolivre/bootcampGo/GoWeb2_3/ProjetoMeli/cmd/server/handler"
	"github.com/brenda-mercadolivre/bootcampGo/GoWeb2_3/ProjetoMeli/internal/users"
	"github.com/brenda-mercadolivre/bootcampGo/GoWeb2_3/ProjetoMeli/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	store := store.Factory("file", "users.json")
	if store == nil {
		log.Fatal("Não foi possivel criar a store")
	}

	repository := users.NewRepository(store)
	service := users.NewService(repository)
	user := handler.NewUser(service)

	router := gin.Default()
	users := router.Group("/users")
	users.POST("/post", user.Create())
	users.GET("/get", user.GetAll())
	users.PUT("/:id", user.Update())
	users.PATCH("/:id", user.UpdateSurnameAndAge())
	users.DELETE("/:id", user.Delete())

	router.Run()
}
