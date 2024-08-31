package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/sudipidus/pismo-test/db"
	"github.com/sudipidus/pismo-test/errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/sudipidus/pismo-test/logger"
	"github.com/sudipidus/pismo-test/services"
)

const ErrorMessageInternalServerError = "Something went wrong"

//var pismoService services.PismoService
//
//func init() {
//	pismoService = &services.PismoServiceImpl{}
//
//}

// @Summary Greetings from Pismo-Test
// @Description Greetings from Pismo-Test
// @Tags greeting/health-check
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Router / [get]
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings from pismo test")
}

// @Summary Create a new account
// @Description Create a new account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param   request  body     services.CreateAccountRequest  true  "Create Account Request"
// @Success 201 {string} string
// @Router /accounts [post]
func AccountsHandler(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Info("creating account")
	//todo: practice an authentication middleware (also endpoint logging, time it)
	// todo: validate create account (operation type and amount whether it's positive or negative)
	var createAccountRequest services.CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&createAccountRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(createAccountRequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}
	response, serviceErr := services.NewPismoService(db.GetStorage()).CreateAccount(r.Context(), createAccountRequest)
	if serviceErr != nil {
		translateErrorAndReturn(w, serviceErr)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{
		Data:    response,
		Success: true,
	}); err != nil {
		// handle the error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	response, serviceErr := services.NewPismoService(db.GetStorage()).FetchAccount(r.Context(), id)
	if serviceErr != nil {
		translateErrorAndReturn(w, serviceErr)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{
		Data:    response,
		Success: true,
	}); err != nil {
		// handle the error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Create a new transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept  json
// @Param   request  body     services.CreateTransactionRequest  true  "Create Transaction Request"
// @Produce  json
// @Success 201 {string} string
// @Router /transactions [post]
func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	var createTransactionRequest services.CreateTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&createTransactionRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(createTransactionRequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	response, serviceErr := services.NewPismoService(db.GetStorage()).CreateTransaction(r.Context(), createTransactionRequest)
	if serviceErr != nil {
		logger.GetLogger().Error(serviceErr.Error())
		translateErrorAndReturn(w, serviceErr)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{
		Data:    response,
		Success: true,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func translateErrorAndReturn(w http.ResponseWriter, err *errors.Error) {
	var statusCode int
	errorMessage := err.Error()

	if err.Code >= 400 && err.Code < 500 {
		statusCode = http.StatusBadRequest
	} else {
		statusCode = http.StatusInternalServerError
		errorMessage = ErrorMessageInternalServerError
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	response := ErrorResponse{
		Code:    statusCode,
		Message: errorMessage,
	}

	// todo: what to suppress and what not to
	jsonResponse, _ := json.Marshal(response)
	_, _ = w.Write(jsonResponse)
}
