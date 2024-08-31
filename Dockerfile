FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download


RUN go build -o pismo main.go

FROM scratch

COPY --from=builder /app/pismo /app/pismo
COPY --from=builder /app/.env /app/.env
COPY --from=builder /app/db/migrations/ /app/db/migrations
COPY --from=builder /app/docs/swagger.json /app/docs/swagger.json

WORKDIR /app

EXPOSE 8080

CMD ["./pismo"]


