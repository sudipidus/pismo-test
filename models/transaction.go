package models

import "time"

type Transaction struct {
	TransactionID   int       `db:"id" json:"transaction_id"`
	AccountID       int       `db:"account_id" json:"account_id"`
	OperationTypeID int       `db:"operation_type_id" json:"operation_type_id"`
	Amount          float64   `db:"amount" json:"amount"`
	Balance         float64   `db:"balance" json:"balance"`
	TransactionDate time.Time `db:"transaction_date" json:"transaction_date"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
