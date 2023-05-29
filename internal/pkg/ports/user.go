package ports

import "github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"

type UserRepository interface {
	Create(value interface{}) error
	First(auth interface{}, conditions ...interface{}) error
}

// drive para user Service debe tenrlos misma frima
type UserService interface {
	Create(user *entity.User) error
	Login(defaultCredentials *entity.DefaultCredential) (string, error)
}
