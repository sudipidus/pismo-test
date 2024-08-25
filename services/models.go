package services

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" example:"1234567890" validate:"required"`
}

type CreateTransactionRequest struct {
	AccountID       int     `json:"account_id" validate:"required" example:"1"`
	OperationTypeID int     `json:"operation_type_id" validate:"required" example:"4"`
	Amount          float64 `json:"amount" validate:"required" example:"123.45"`
}
