package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
