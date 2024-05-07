package main

import "github.com/shopspring/decimal"

func incrementBalance(acc *Account, amount string) {
	amountDecimal, _ := decimal.NewFromString(amount)

	acc.InitialBalance = acc.InitialBalance.Add(amountDecimal)
}

func decrementBalance(acc *Account, amount string) {
	amountDecimal, _ := decimal.NewFromString(amount)

	acc.InitialBalance = acc.InitialBalance.Sub(amountDecimal)
}
