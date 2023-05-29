package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/ports"
)

type userHandler struct {
	userService ports.UserService
}

func newHandler(service ports.UserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}

// <Metdos paat contesta las peticiones HTTP
func (u *userHandler) CreateUser(context *gin.Context) {

	user := &entity.User{}

	//bidinges eecir recibir la inforacion
	//tomamos el cuepro de la t¿peticon
	if err := context.Bind(user); err != nil {
		context.JSON(http.StatusBadRequest, errors.New("Invalid Input"))
		return
	}

	if err := u.userService.Create(user); err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	user.Password = ""
	context.JSON(http.StatusCreated, user)

}

func (u *userHandler) Login(contex *gin.Context) {

	credentials := &entity.DefaultCredential{}

	if err := contex.Bind(credentials); err != nil {
		contex.JSON(http.StatusBadRequest, errors.New("Invalid Credentials"))
		return
	}

	token, err := u.userService.Login(credentials)

	if err != nil {
		contex.JSON(http.StatusUnauthorized, nil)
		return
	}

	contex.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
