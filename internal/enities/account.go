package enities

import (
	"time"
)

type Account struct {
	Id int
	Name string
	CreateAt time.Time
	UpdateAt time.Time
	Currency string
	InitialBalance float64
}
