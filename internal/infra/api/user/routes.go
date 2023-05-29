package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/infra/repositories/postgres"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/service/user"
)

func RegisterRoutes(engine *gin.Engine) {

	//SE PUEDE EN UNA FUNCION INIT?
	userRepo := postgres.NewClient()
	userService := user.NewService(userRepo)
	handler := newHandler(userService)
	//INIT DETRO HANDLER -LE ASIGNAMOS VARIABLES GLOABLES

	engine.POST("/api/v1/user", handler.CreateUser)

	//login
	engine.POST("/api/v1/auth", handler.Login)
}
