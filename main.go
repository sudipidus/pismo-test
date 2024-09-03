package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sudipidus/pismo-test/config"
	"github.com/sudipidus/pismo-test/db"
	_ "github.com/sudipidus/pismo-test/db"
	"github.com/sudipidus/pismo-test/logger"
	"github.com/sudipidus/pismo-test/routes"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	r.Use(loggingMiddleware)

	// need to explicitly map because it's gorilla/mux (blank import registers handlers for native mux)
	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
		<-sigterm
		logger.GetLogger().Info("Received SIGTERM, shutting down...")
		cancel()
	}()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		logger.GetLogger().Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()
	fmt.Println("Server listening on port 8080, visit http://localhost:8080/swagger/index.html")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Request %s %s from %s",
			r.Method, r.RequestURI, r.RemoteAddr)
		logger.GetLogger().Info(msg)
		next.ServeHTTP(w, r)
	})
}
