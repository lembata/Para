package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lembata/para/internal/enities"
	"github.com/lembata/para/pkg/database"
)

type AccountService struct {
	templates Templates
}

type ApiResponse struct {
	Success   bool   `json:"success"`
	Data      any    `json:"data"`
	Error     string `json:"error"`
	ErrorCode int    `json:"errorCode"`
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
}

func (s *AccountService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(2 * time.Second)

	var account AccountData
	err := json.NewDecoder(r.Body).Decode(&account)

	logger.Debugf("Creating account: %v", account)

	if account.AccountName == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := Failure("account name is required", http.StatusBadRequest)
		w.Write(response)
		return
	}

	if len(account.Currency) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("currency is invalid"))
		return
	}

	//newAccount := enities.Account{
	_ = enities.Account{
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
	}

	db := database.NewDatabase()

	err = db.Open("")
	defer db.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Failure("failed to open database", http.StatusInternalServerError))
		return
	}
}

func (s *AccountService) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.RouteContext(ctx).URLParam("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Failure("invalid account id", http.StatusBadRequest))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(Data(AccountData {
		Id: id,
		AccountName: "Test Account",
		Currency: "EUR",
		IBAN: "DE1234567890",
		BIC: "BIC1234567890",
		AccountNumber: "1234567890",
		OpeningBalance: 100.00,
		OpeningBalanceDate: "2021-01-01",
		Notes: "This is a test account",
	}))
}

func Failure(error string, errorCode int) []byte {
	json, _ := json.Marshal(ApiResponse{
		Success:   false,
		Error:     error,
		ErrorCode: errorCode,
	})

	return json
}

func Success() []byte {
	json, _ := json.Marshal(ApiResponse{
		Success:   true,
	})

	return json
}

func Data(data any) []byte {
	json, err := json.Marshal(ApiResponse{
		Success:   true,
		Data:      data,
	})

	if err != nil {
		return Failure("failed to serialize data", http.StatusInternalServerError)
	}

	return json
}
