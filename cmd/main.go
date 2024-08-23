package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/sudipidus/pismo-test/docs"
	"github.com/sudipidus/pismo-test/routes"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	routes.SetupRoutes(r)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
