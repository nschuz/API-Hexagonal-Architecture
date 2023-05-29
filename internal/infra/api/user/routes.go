package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/infra/repositories/postgres"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/service/user"
)

func RegisterRoutes(engine *gin.Engine) {

	//tos s epuede crera en unaa funcion init
	userRepo := postgres.NewClient()
	userService := user.NewService(userRepo)
	handler := newHandler(userService)
	//init deentr handler y lo asignamso a avriables globales

	engine.POST("/api/v1/user", handler.CreateUser)

	//login
	engine.POST("/api/v1/auth", handler.Login)
}
