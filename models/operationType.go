package models

import "time"

type OperationType struct {
	ID          int       `db:"id" json:"id"`
	Type        string    `db:"type" json:"type"`
	Description string    `db:"description" json:"description"`
	IsCredit    bool      `db:"is_credit" json:"debit_credit"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
