// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID          int64
	Balance     int64
	Name        string
	PhoneNumber string
	Currency    string
	CreatedAt   time.Time
}

type Transaction struct {
	ID         int64
	RecieverID int64
	SenderID   int64
	Amount     int64
	Currency   string
	Message    sql.NullString
	Deadline   sql.NullTime
	CreatedAt  time.Time
}