package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brenda-mercadolivre/bootcampGo/GoWeb2_3/ProjetoMeli/cmd/server/handler"
	"github.com/brenda-mercadolivre/bootcampGo/GoWeb2_3/ProjetoMeli/internal/users"
	"github.com/brenda-mercadolivre/bootcampGo/GoWeb2_3/ProjetoMeli/pkg/store"
	"github.com/brenda-mercadolivre/bootcampGo/GoWeb2_3/ProjetoMeli/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func respondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, web.NewResponse(code, nil, message))
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("por favor, configure a variável de ambiente - token")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, http.StatusUnauthorized, "Token API obrigatório")
			return
		}

		if token != requiredToken {

			respondWithError(c, http.StatusUnauthorized, "Token API inválido")
			return
		}
		c.Next()
	}
}
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
	users.Use(TokenAuthMiddleware())
	users.POST("/post", user.Create())
	users.GET("/get", user.GetAll())
	users.PUT("/:id", user.Update())
	users.PATCH("/:id", user.UpdateSurnameAndAge())
	users.DELETE("/:id", user.Delete())

	router.Run()
}
