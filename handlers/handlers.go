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
	logger.GetLogger().Info("creating account")
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

	json.NewEncoder(w).Encode(Response{
		Data:    response,
		Success: true,
	})
}

func translateErrorAndReturn(w http.ResponseWriter, err *errors.Error) {
	if err.Code >= 500 {
		w.WriteHeader(http.StatusInternalServerError)
	} else if err.Code >= 400 && err.Code < 500 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
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

	json.NewEncoder(w).Encode(Response{
		Data:    response,
		Success: true,
	})
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
		translateErrorAndReturn(w, serviceErr)
		return
	}

	json.NewEncoder(w).Encode(Response{
		Data:    response,
		Success: true,
	})
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
}
