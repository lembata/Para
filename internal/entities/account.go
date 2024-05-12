package entities

import (
	"time"
)

type AccountEntity struct {
	Id                 int64
	Name               string
	CreateAt           time.Time
	UpdateAt           time.Time
	Currency           string
	IBAN               string
	BIC                string
	AccountNumber      string
	OpeningBalance     int
	OpeningBalanceDate string
	Notes              string
}
