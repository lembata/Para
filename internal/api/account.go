package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lembata/para/internal/entities"
	"github.com/lembata/para/pkg/database"
)

type AccountService struct {
	templates Templates
}

type AccountData struct {
	Id                 int     `json:"id"`
	AccountName        string  `json:"accountName"`
	Currency           string  `json:"currency"`
	IBAN               string  `json:"iban"`
	BIC                string  `json:"bic"`
	AccountNumber      string  `json:"accountNumber"`
	OpeningBalance     float64 `json:"openingBalance"`
	OpeningBalanceDate string  `json:"openiningBalanceDate"`
	Notes              string  `json:"notes"`
	IncludeInNetWorth  bool    `json:"includeInNetWorth"`
}

type AccountShort struct {
	Id             int     `json:"id"`
	AccountName    string  `json:"accountName"`
	CurrentBalance float64 `json:"currentBalance"`
	LastActivity   float64 `json:"currentBalance"`
}

func (s *AccountService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account AccountData
	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Debugf("Creating account: %v", account)

	if account.AccountName == "" {
		WriteFailure(w, "account name is required", http.StatusBadRequest)
		return
	}

	if len(account.Currency) != 3 {
		WriteFailure(w, "currency is invalid", http.StatusBadRequest)
		return
	}

	newAccount := entities.AccountEntity{
		//_ = entities.AccountEntity{
		Id:                 0,
		Name:               account.AccountName,
		CreateAt:           time.Now(),
		UpdateAt:           time.Now(),
		Currency:           account.Currency,
		OpeningBalance:     int(account.OpeningBalance * 100),
		OpeningBalanceDate: account.OpeningBalanceDate,
		IBAN:               account.IBAN,
		BIC:                account.BIC,
		AccountNumber:      account.AccountNumber,
		Notes:              account.Notes,
		IncludeInNetWorth:  account.IncludeInNetWorth,
	}

	db := database.GetInstance()
	ctx, err := db.Begin(r.Context(), false)

	if err != nil {
		logger.Errorf("failed to begin transaction: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.CreateAccount(ctx, newAccount)

	if err != nil {
		logger.Errorf("failed to create account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	time.Sleep(200 * time.Second)

	err = db.Commit(ctx)

	if err != nil {
		logger.Errorf("failed to create account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(Success())
}

func (s *AccountService) All(w http.ResponseWriter, r *http.Request) {
	var tableRequest TableRequest
	err := json.NewDecoder(r.Body).Decode(&tableRequest)

	if err != nil {
		logger.Errorf("failed to get account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.GetInstance()

	//ctx, err := db.Begin(r.Context(), false)
	var ctx context.Context

	if ctx, err = db.Begin(r.Context(), false); err != nil {
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	var rows *sql.Rows
	if rows, err = db.GetAccounts(ctx); err != nil {
		logger.Errorf("failed to get account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if err = db.Commit(ctx); err != nil {
	// 	WriteFailure(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	
	logger.Debug("HERE 0")

	for rows.Next() {
        //var alb Album
		var id int64
		var name string
		var delta int64

        if err := rows.Scan(&id, &name, &delta); err != nil {
            logger.Errorf("Error %v", err)
			break
        }

	logger.Debugf("id: %d, name: %s, delta: %d" , id, name, delta)
    }
		
	logger.Debug("HERE 1")


	if err = db.Commit(ctx); err != nil {
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	return
}

func (s *AccountService) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := chi.RouteContext(r.Context())
	ctx.URLParam("id")
	idStr := ctx.URLParam("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Failure("invalid account id", http.StatusBadRequest))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(Data(AccountData{
		Id:                 id,
		AccountName:        "Test Account",
		Currency:           "EUR",
		IBAN:               "DE1234567890",
		BIC:                "BIC1234567890",
		AccountNumber:      "1234567890",
		OpeningBalance:     100.00,
		OpeningBalanceDate: "2021-01-01",
		Notes:              "This is a test account",
	}))
}
