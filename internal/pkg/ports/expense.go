package ports

import (
	"time"

	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
)

type ExpeseRepository interface {
	AddExpense(expense *entity.Expense) error
	GetUserExpenses(userID int, startDate time.Time, endDate time.Time) ([]*entity.Expense, error)
}
