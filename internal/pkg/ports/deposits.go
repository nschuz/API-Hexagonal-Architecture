package ports

import (
	"time"

	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
)

type DespositRepository interface {
	AddDeposit(expense *entity.Deposit) error
	GetUserDeposits(userID int, startDate time.Time, endDate time.Time) ([]*entity.Deposit, error)
}

type DepositService interface {
	//metodos que estan en nuestro servicio
	Add(deposit *entity.Deposit) error
	GetUserDeposits(userID int, startDate, endDate string) ([]*entity.Deposit, error)
}
