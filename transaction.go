package main

import (
	"errors"
	"net/http"

	"github.com/shopspring/decimal"
)

type TransactionRequest struct {
	SourceAccountID      decimal.Decimal `json:"source_account_id"`
	DestinationAccountID decimal.Decimal `json:"destination_account_id"`
	Amount               string          `json:"amount"`
}

func transfer(repo *Repo, sourceAccountId decimal.Decimal, destinationAccountId decimal.Decimal, amount string) {
	sourceAccountIdStr := sourceAccountId.String()
	destinationAccountIdStr := destinationAccountId.String()

	sourceAccount, _ := repo.GetAccount(sourceAccountIdStr)
	destinationAccoun, _ := repo.GetAccount(destinationAccountIdStr)
	decrementBalance(sourceAccount, amount)
	incrementBalance(destinationAccoun, amount)
}

func (a *TransactionRequest) Bind(r *http.Request) error {
	if a.Amount == "" {
		return errors.New("missing required Transaction fields.")
	}

	return nil
}
