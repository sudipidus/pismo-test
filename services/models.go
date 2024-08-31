package services

type CreateAccountRequest struct {
	//todo: add more validation
	DocumentNumber string `json:"document_number" example:"1234567890" validate:"required"`
}

type CreateTransactionRequest struct {
	AccountID       int     `json:"account_id" validate:"required" example:"1"`
	OperationTypeID int     `json:"operation_type_id" validate:"required,oneof=1 2 3 4" example:"4"`
	Amount          float64 `json:"amount" validate:"required" example:"123.45"`
}
