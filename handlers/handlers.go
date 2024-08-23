package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//todo: add separate handlers

// @Summary Greetings from Pismo-Test
// @Description Greetings from Pismo-Test
// @Tags root
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Router / [get]
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings from Pismo-Test")
}

// @Summary Create a new account
// @Description Create a new account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 201 {string} string
// @Router /accounts [post]
func AccountsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Account has been created")
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
// @Produce  json
// @Success 201 {string} string
// @Router /transactions [post]
func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "new transaction created")
}
