package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	repo := NewRepo()

	router.Post("/accounts", CreateAccount(repo))
	router.Get("/accounts/{account_id}", GetAccount(repo))
	router.Post("/transactions", CreateTransaction(repo))

	http.ListenAndServe(":3000", router)
}

func CreateAccount(repo *Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account := &AccountRequest{}

		if err := render.Bind(r, account); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		acc, _ := repo.CreateAccount(account)

		render.Status(r, http.StatusCreated)
		render.Render(w, r, acc)
	}
}

func GetAccount(repo *Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountID := chi.URLParam(r, "account_id")

		acc, _ := repo.GetAccount(accountID)

		render.Status(r, http.StatusOK)
		render.Render(w, r, acc)
	}
}

func CreateTransaction(repo *Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transaction := &TransactionRequest{}

		if err := render.Bind(r, transaction); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		transfer(repo, transaction.SourceAccountID, transaction.DestinationAccountID, transaction.Amount)

		render.Status(r, http.StatusCreated)
	}
}
