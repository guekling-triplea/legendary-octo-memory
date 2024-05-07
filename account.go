package main

import (
	"errors"
	"net/http"

	"github.com/shopspring/decimal"
)

type AccountRequest struct {
	AccountID      decimal.Decimal `json:"account_id"`
	InitialBalance string          `json:"initial_balance"`
}

type Account struct {
	AccountID      string          `json:"account_id"`
	InitialBalance decimal.Decimal `json:"initial_balance"`
}

func (a *Account) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *AccountRequest) Bind(r *http.Request) error {
	if a.InitialBalance == "" {
		return errors.New("missing required Account fields.")
	}

	return nil
}
