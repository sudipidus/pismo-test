### Pismo Transaction Service Demo

Dependencies:
1. go1.23
2. golang-migrate for schema migration `github.com/golang-migrate/migrate/v4/cmd/migrate`
3. swaggo for documentation generation `github.com/swaggo/swag/cmd/swag`
4. postgres DB
5. gomock for mock generation
6. godotevn for env management `github.com/joho/godotenv`

#### Starting Up
tldr -> run the start script `sh start.sh`
1. Clone the repo
2. Install go1.23 (use a version manager like gvm: `gvm install go1.23 && gvm use go1.23`)
3. Download the dependencies (`go mod download`)
4. Create mocks and documentation (`go generate ./... && swag init`) (or run `sh generate_doc.sh`)
5. Start the dependencies `docker compose up db`
6. Start the app `go run ./...`
7. Access the swagger UI `http://localhost:8080/swagger/index.html`

![swagger-ui](swagger-ui.png)

#### Dockerized Way
`docker compose up`
(This has all the dependencies as well as migration setup)


#### Testing
This has 3 layers of testing:
- handler (controller testing)
- service layer testing
- storage (postgres) testing

Setup up test postgres `docker compose up db-test`

`go test ./...`