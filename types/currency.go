package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

var ErrInvalidCurrency = fmt.Errorf("invalid currency")

type Currency string

const (
	BRL Currency = "BRL"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

func NewCurrency(value string) (Currency, error) {
	c := Currency(strings.ToUpper(value))
	if !c.IsValid() {
		return "", ErrInvalidCurrency
	}
	return c, nil
}

func (c Currency) String() string {
	return string(c)
}

func (c Currency) IsValid() bool {
	switch c {
	case BRL, USD, EUR: // CORREÇÃO: Adicionado EUR à lista de moedas válidas.
		return true
	default:
		return false
	}
}

func (c Currency) IsEmpty() bool {
	return c == ""
}

func (c Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Currency) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*c = Currency(s)
	return nil
}

func (c Currency) Value() (driver.Value, error) {
	if c.IsEmpty() {
		return nil, nil
	}
	return c.String(), nil
}

func (c *Currency) Scan(src interface{}) error {
	if src == nil {
		*c = ""
		return nil
	}

	var s string
	switch v := src.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return ErrInvalidCurrency
	}

	*c = Currency(s)
	return nil
}
