// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: accounts.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO Accounts (
  name, phone_number,currency,balance
) VALUES (
  $1, $2, $3, $4
) RETURNING id, name, balance, phone_number, currency, created_at
`

type CreateAccountParams struct {
	Name        string
	PhoneNumber string
	Currency    string
	Balance     int64
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.Name,
		arg.PhoneNumber,
		arg.Currency,
		arg.Balance,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.PhoneNumber,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, name, balance, phone_number, currency, created_at FROM Accounts 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.PhoneNumber,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const updateAccountBalance = `-- name: UpdateAccountBalance :exec
UPDATE Accounts
SET balance = $2
WHERE id = $1
`

type UpdateAccountBalanceParams struct {
	ID      int64
	Balance int64
}

func (q *Queries) UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountBalance, arg.ID, arg.Balance)
	return err
}

const updateAccountPhoneNumber = `-- name: UpdateAccountPhoneNumber :exec
UPDATE Accounts
SET phone_number = $2
WHERE id = $1
`

type UpdateAccountPhoneNumberParams struct {
	ID          int64
	PhoneNumber string
}

func (q *Queries) UpdateAccountPhoneNumber(ctx context.Context, arg UpdateAccountPhoneNumberParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountPhoneNumber, arg.ID, arg.PhoneNumber)
	return err
}
