package expense

import (
	"time"

	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/ports"
	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/utils"
)

type service struct {
	repo ports.ExpeseRepository
}

func NewService(repo ports.ExpeseRepository) *service {
	return &service{repo: repo}
}

func (s *service) Add(expense *entity.Expense) error {
	return s.repo.AddExpense(expense)
}

func (s *service) GetUserExpenses(userID int, startDate, endDate string) ([]*entity.Expense, error) {

	var st, ed time.Time
	var err error

	if !utils.IsDateRangeValid(startDate, endDate) {
		st, ed = utils.GetLastSevenDays()
	} else {
		st, ed, err = utils.ParseDateRange(startDate, endDate)
		if err != nil {
			return nil, err
		}
	}

	return s.repo.GetUserExpenses(userID, st, ed)

}
