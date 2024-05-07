package main

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

type Repo struct {
	storage sync.Map
}

// NewRepo returns a new in memory repo
func NewRepo() *Repo {
	return &Repo{
		storage: sync.Map{},
	}
}

func (r *Repo) CreateAccount(cap *AccountRequest) (*Account, error) {
	initialBalance, _ := decimal.NewFromString(cap.InitialBalance)
	accountID := cap.AccountID.String()

	// create a new account
	acc := &Account{
		AccountID:      accountID,
		InitialBalance: initialBalance,
	}

	// store the account in the storage map
	r.storage.Store(acc.AccountID, acc)

	return acc, nil
}

func (r *Repo) GetAccount(id string) (*Account, error) {
	// get the account from the storage map
	acc, okLoad := r.storage.Load(id)
	if !okLoad {
		return nil, fmt.Errorf("Account not found")
	}

	accTyped, ok := acc.(*Account) //nolint:go-critic // type assertion is valid
	if !ok {
		return nil, fmt.Errorf("Unexpected Type Error")
	}

	return accTyped, nil
}
