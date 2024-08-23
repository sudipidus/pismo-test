package routes

import (
	"github.com/gorilla/mux"
	"github.com/sudipidus/pismo-test/handlers"
)

func SetupRoutes(r *mux.Router) {
	// Define routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/accounts", handlers.AccountsHandler).Methods("POST")
	r.HandleFunc("/accounts/{accountID}", handlers.GetAccountsHandler).Methods("GET")
	r.HandleFunc("/transactions", handlers.TransactionHandler).Methods("POST")
}
