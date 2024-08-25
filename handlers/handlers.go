package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/sudipidus/pismo-test/services"
)

var pismoService services.PismoService

func init() {
	pismoService = &services.PismoServiceImpl{}

}

// @Summary Greetings from Pismo-Test
// @Description Greetings from Pismo-Test
// @Tags greeting/health-check
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Router / [get]
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings")
}

// @Summary Create a new account
// @Description Create a new account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param   request  body     CreateAccountRequest  true  "Create Account Request"
// @Success 201 {string} string
// @Router /accounts [post]
func AccountsHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Account has been created with document number: "+req.DocumentNumber)
}

// @Summary Get an account by ID
// @Description Get an account by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param   accountID  path     string  true  "Account ID"
// @Success 200 {string} string
// @Router /accounts/{accountID} [get]
func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["accountID"]
	fmt.Fprintf(w, "Account with ID %s", id)
}

// @Summary Create a new transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept  json
// @Param   request  body     CreateTransactionRequest  true  "Create Transaction Request"
// @Produce  json
// @Success 201 {string} string
// @Router /transactions [post]
func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "new transaction created")
}

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" example:"1234567890" validate:"required"`
}

type CreateTransactionRequest struct {
	AccountID       int     `json:"account_id" validate:"required" example:"1"`
	OperationTypeID int     `json:"operation_type_id" validate:"required" example:"4"`
	Amount          float64 `json:"amount" validate:"required" example:"123.45"`
}
