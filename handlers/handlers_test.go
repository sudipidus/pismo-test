package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sudipidus/pismo-test/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func init() {
//	logger.InitLogger()
//	os.Setenv("DB_DSN", "postgres://pismo-user:pismo-secret@db-test:5433/pismo?sslmode=disable")
//	db.Init()
//}

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlers.HomeHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "Greetings from pismo test" {
		t.Errorf("expected response body 'Greetings from pismo test', got %s", w.Body.String())
	}
}

func TestAccountsHandler(t *testing.T) {
	reqBody := `{"document_number":"1234567890"}`
	req, err := http.NewRequest("POST", "/accounts", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.AccountsHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
	var response struct {
		Data    interface{} `json:"data"`
		Success bool        `json:"success"`
	}
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	if !response.Success {
		t.Errorf("expected success to be true, got %v", response.Success)
	}
}

func TestAccountsHandler_ValidationError(t *testing.T) {
	reqBody := `{"name": "John Doe"}`
	req, err := http.NewRequest("POST", "/accounts", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.AccountsHandler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetAccountsHandler(t *testing.T) {
	reqBody := `{"document_number":"1234567890"}`
	_, err := http.NewRequest("POST", "/accounts", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	//todo: test with the returned ID instead, now since it's serial PK 1 works

	req, err := http.NewRequest("GET", "/accounts/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	req = mux.SetURLVars(req, map[string]string{"accountID": "1"})
	handlers.GetAccountsHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
	var response struct {
		Data    interface{} `json:"data"`
		Success bool        `json:"success"`
	}
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	if !response.Success {
		t.Errorf("expected success to be true, got %v", response.Success)
	}
}

func TestTransactionHandler(t *testing.T) {
	reqBody := `{"document_number":"1234567890"}`
	_, err := http.NewRequest("POST", "/accounts", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	//todo: create transaction with returned account_id
	reqBody = `{"account_id": 1, "amount": 10.99, "operation_type_id":1}`
	req, err := http.NewRequest("POST", "/transactions", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.TransactionHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
	var response struct {
		Data    interface{} `json:"data"`
		Success bool        `json:"success"`
	}
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	if !response.Success {
		t.Errorf("expected success to be true, got %v", response.Success)
	}
}
