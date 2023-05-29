package deposits

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/ports"
)

type depositsHandler struct {
	depositService ports.DepositService
}

func newHandler(service ports.DepositService) *depositsHandler {
	return &depositsHandler{
		depositService: service,
	}
}

func (d *depositsHandler) Add(context *gin.Context) {

	//BIND PARA LLER EL CUPERPO DE LA PETICION
	deposit := &entity.Deposit{}

	if err := context.Bind(deposit); err != nil {
		context.JSON(http.StatusBadRequest, errors.New("Invalid deposit fromat"))
		return
	}

	if err := d.depositService.Add(deposit); err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	context.JSON(http.StatusCreated, nil)

}
