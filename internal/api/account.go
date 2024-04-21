package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/hcl/hcl/strconv"
	"github.com/lembata/para/internal/enities"
	"github.com/lembata/para/pkg/database"
)

type AccountService struct {
	templates Templates
}

func (s *AccountService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	initialAmount, err := strconv.ParseFloat(r.Form.Get("initial-balance", 64))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid initial balance"))
		return
	}

	newAccount := enities.Account{
		Id: 0,
		Name: r.Form.Get("name"),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		Currency: r.Form.Get("currency"),
		InitialBalance: initialAmount,
	}

	db := database.NewDatabase();

	err = db.Open("")
	defer db.Close()

	if  err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}



}

