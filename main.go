package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sudipidus/pismo-test/config"
	"github.com/sudipidus/pismo-test/db"
	_ "github.com/sudipidus/pismo-test/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sudipidus/pismo-test/logger"
	"github.com/sudipidus/pismo-test/routes"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
)

// @title Pismo Transaction Service - Demo
// @version 1.0
// @description This is a simplified transaction service.
// @termsOfService http://swagger.io/terms/

// @contact.name Sudip Bhandari
// @contact.url https://sudipidus.github.io
// @contact.email sudip.post@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	logger.InitLogger()

	config.Init()

	db.Init()

	db.SeedOperationType(db.GetStorage())

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Err loading .env file: %v", err)
	}

	r := mux.NewRouter()

	routes.SetupRoutes(r)

	r.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
