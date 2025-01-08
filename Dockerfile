FROM golang:1.23-alpine3.20

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

CMD ["sh", "-c", "go run ./internal/migrate/migrate.go && go run ./cmd/main.go"]
