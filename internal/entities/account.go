package entities

import (
	"time"
)

type AccountEntity struct {
	Id                 int64     `db:"id" json:"id"`
	Name               string    `db:"name" json:"name"`
	CreateAt           time.Time `db:"create_at" json:"createAt"`
	UpdateAt           time.Time `db:"update_at" json:"updateAt"`
	Currency           string    `db:"currency" json:"currency"`
	IBAN               string    `db:"iban" json:"iban"`
	BIC                string    `db:"bic" json:"bic"`
	AccountNumber      string    `db:"account_number" json:"accountNumber"`
	OpeningBalance     int       `db:"opening_balance" json:"openingBalance"`
	OpeningBalanceDate string    `db:"opening_balance_date" json:"openingBalanceDate"`
	Notes              string    `db:"notes" json:"notes"`
	IncludeInNetWorth  bool      `db:"include_in_net_worth" json:"includeInNetWorth"`
}

type AccountRow struct {
	Id      int           `json:"id"`
	Name    string        `json:"name"`
	Balance CurrencyValue `json:"balance"`
}

type CurrencyValue struct {
	Currency string `json:"currency"`
	Value    int    `json:"value"`
}
