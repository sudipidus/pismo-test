#!/bin/bash
set -e

echo "Downloading Go modules..."
go mod download

echo "Installing swag CLI tool..."
go install github.com/swaggo/swag/cmd/swag@latest

echo "Running Go generate and initializing swag..."
go generate ./...

echo "Starting the database service in the background using Docker Compose..."
docker compose up -d db

echo "Installing golang-migrate CLI tool..."
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

echo "Running database migrations..."
migrate -database "postgres://pismo-user:pismo-secret@localhost:5433/pismo?sslmode=disable" -path db/migrations up

echo "Running the Go application..."
go run ./...
