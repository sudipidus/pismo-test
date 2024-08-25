package models

import "time"

type Account struct {
	ID             int       `db:"id" json:"id"`
	DocumentNumber string    `db:"document_number" json:"document_number"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
