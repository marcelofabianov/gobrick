package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

var ErrInvalidDay = fmt.Errorf("day must be between 1 and 31")

type Day int

func NewDay(value int) (Day, error) {
	if value < 1 || value > 31 {
		return 0, ErrInvalidDay
	}
	return Day(value), nil
}

func (d Day) Int() int {
	return int(d)
}

// HasPassed verifica se este dia já passou no mês corrente.
func (d Day) HasPassed(today time.Time) bool {
	return d.Int() < today.Day()
}

// DaysUntil calcula quantos dias faltam para a próxima ocorrência deste dia.
func (d Day) DaysUntil(today time.Time) int {
	day := d.Int()
	todayDay := today.Day()

	if day >= todayDay {
		return day - todayDay
	}

	daysInMonth := time.Date(today.Year(), today.Month()+1, 0, 0, 0, 0, 0, today.Location()).Day()
	return (daysInMonth - todayDay) + day
}

// DaysOverdue calcula quantos dias se passaram desde a última ocorrência deste dia.
func (d Day) DaysOverdue(today time.Time) int {
	day := d.Int()
	todayDay := today.Day()

	if day <= todayDay {
		return todayDay - day
	}

	prevMonth := today.AddDate(0, -1, 0)
	daysInPrevMonth := time.Date(prevMonth.Year(), prevMonth.Month()+1, 0, 0, 0, 0, 0, today.Location()).Day()
	return (daysInPrevMonth - day) + todayDay
}

func (d Day) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Int())
}

func (d *Day) UnmarshalJSON(data []byte) error {
	var day int
	if err := json.Unmarshal(data, &day); err != nil {
		return err
	}
	*d = Day(day)
	return nil
}

func (d Day) Value() (driver.Value, error) {
	return int64(d.Int()), nil
}

func (d *Day) Scan(src interface{}) error {
	if src == nil {
		*d = 0
		return nil
	}

	var day int64
	switch v := src.(type) {
	case int64:
		day = v
	default:
		return fmt.Errorf("unsupported scan type for Day: %T", src)
	}

	*d = Day(day)
	return nil
}
