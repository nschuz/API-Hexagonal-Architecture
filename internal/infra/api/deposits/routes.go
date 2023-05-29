package deposits

import (
	"github.com/gin-gonic/gin"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/infra/api/middlewares"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/infra/repositories/firestore"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/service/deposits"
)

func RegisterRoutes(e *gin.Engine) {
	repo := firestore.NewClient()
	service := deposits.NewService(repo)
	handler := newHandler(service)

	//primero ejecuta el middlearw y despues el handler
	e.POST("/api/v1/deposits", middlewares.Authenticate(), handler.Add)

}
