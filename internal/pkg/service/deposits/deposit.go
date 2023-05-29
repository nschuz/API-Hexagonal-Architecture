package deposits

import (
	"time"

	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/ports"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/utils"
)

type service struct {
	repo ports.DespositRepository
}

func NewService(repo ports.DespositRepository) *service {
	return &service{repo: repo}
}

// METODO
func (s *service) Add(deposit *entity.Deposit) error {
	return s.repo.AddDeposit(deposit)
}

//es porparel la lgogica de neogico obtener la fecha en strings

func (s *service) GetUserDeposits(userID int, startDate, endDate string) ([]*entity.Deposit, error) {

	var st, ed time.Time
	var err error
	//Verificar que le inervalo de fecha sea correcto
	if !utils.IsDateRangeValid(startDate, endDate) {
		//gernamos un rango de fechas por default
		st, ed = utils.GetLastSevenDays()
	} else {
		st, ed, err = utils.ParseDateRange(startDate, endDate)
	}

	deposits, err := s.repo.GetUserDeposits(userID, st, ed)
	if err != nil {
		return nil, err
	}

	return deposits, nil

}
