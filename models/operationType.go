package models

import "time"

type OperationType struct {
	ID          int       `json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	DebitCredit string    `json:"debit_credit"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
