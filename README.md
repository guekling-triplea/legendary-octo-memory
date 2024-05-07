# legendary-octo-memory

## Running the project

`go run main.go repo.go balance.go account.go transaction.go error.go`

## What Is Not Done

- Importing local packages
- Creating an account with an account_id linked to an existing account MAY overwrite the existing account
- No check of sufficient balance before transaction submission
- No check of whether the `source_account_id` and `destination_account_id` is valid before transaction submission
- Tests