package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/infra/api/deposits"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/infra/api/user"
)

func RegisterRoutes(e *gin.Engine) {
	//USER ROUTES
	user.RegisterRoutes(e)

	//DEPOSIT ROUTES
	deposits.RegisterRoutes(e)

}
