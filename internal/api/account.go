package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lembata/para/internal/entities"
	"github.com/lembata/para/pkg/currency"
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
	LastActivity   string  `json:"lastActivity"`
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
		OpeningBalance:     currency.ToCoins(account.OpeningBalance),
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

	err = db.Commit(ctx)

	if err != nil {
		logger.Errorf("failed to create account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = WriteSuccess(w)
}

func (s *AccountService) EditAccount(w http.ResponseWriter, r *http.Request) {
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
		Id:                 int64(account.Id),
		Name:               account.AccountName,
		UpdateAt:           time.Now(),
		Currency:           account.Currency,
		OpeningBalance:     currency.ToCoins(account.OpeningBalance),
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

	if _, err = db.EditAccount(ctx, newAccount); err != nil {
		logger.Errorf("failed to edit account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Commit(ctx)

	if err != nil {
		logger.Errorf("failed to create account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = WriteSuccess(w)
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

	var accounts []entities.AccountRow

	if accounts, err = db.GetAccounts(ctx, 0, 10, "id"); err != nil {
		logger.Errorf("failed to get account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = db.Commit(ctx); err != nil {
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Debugf("Got accounts: %v", accounts)

	_, _ = WriteData(w, accounts)
}

func (s *AccountService) GetAccount(w http.ResponseWriter, r *http.Request) {
	chiCtx := chi.RouteContext(r.Context())
	chiCtx.URLParam("id")
	idStr := chiCtx.URLParam("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		WriteFailure(w, "invalid account id", http.StatusBadRequest)
		return
	}

	db := database.GetInstance()

	var ctx context.Context
	if ctx, err = db.Begin(r.Context(), false); err != nil {
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	var account entities.AccountEntity

	if account, err = db.GetAccountById(ctx, int64(id)); err != nil {
		logger.Errorf("failed to get account: %v", err)
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = db.Commit(ctx); err != nil {
		WriteFailure(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountData := AccountData{
		Id:                 int(account.Id),
		AccountName:        account.Name,
		Currency:           account.Currency,
		IBAN:               account.IBAN,
		BIC:                account.BIC,
		AccountNumber:      account.AccountNumber,
		OpeningBalance:     currency.FromCoins(account.OpeningBalance),
		OpeningBalanceDate: account.OpeningBalanceDate,
		Notes:              account.Notes,
		IncludeInNetWorth:  account.IncludeInNetWorth,
	}

	_, _ = WriteData(w, accountData)
}
