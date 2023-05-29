package utils

import (
	"fmt"
	"time"
)

func ParseDateRange(startDate, endDate string) (time.Time, time.Time, error) {

	st, err := time.Parse("2006-01-02", startDate)

	if err != nil {
		return time.Now(), time.Now(), err
	}

	ed, err := time.Parse("2006-01-02", startDate)

	ed = time.Date(ed.Year(), ed.Month(), ed.Day(), 23, 59, 59, 1e9-1, ed.Location())
	if err != nil {
		return time.Now(), time.Now(), err
	}

	return st, ed, nil

}

func IsDateRangeValid(startDate, endDate string) bool {
	if startDate == "" || endDate == "" {
		return false
	}

	st, ed, err := ParseDateRange(startDate, endDate)
	if err != nil {
		return false
	}

	return ed.After(st)

}

func GetLastSevenDays() (time.Time, time.Time) {
	today := time.Now().UTC()

	//esta hora cuando entra el servidor pero antes no los vamos a ver
	//esta fecha la debemos redodnear a nuestro startdate para tenr el dia completo
	startDate := today.AddDate(0, 0, -7).Truncate(time.Hour * 24)                                       //ROUND REDONDE A LA HORA CERO MAS CERNCANA
	endDate := time.Date(today.Year(), today.Month(), today.Day(), 23, 59, 59, 1e6-1, today.Location()) //el ultimo momento del dia de hoy

	fmt.Println(startDate) //desde la hora 0
	fmt.Println(endDate)   //ultimo nanosegudno del dia
	return startDate, endDate
}
