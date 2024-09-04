FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .env ./

COPY . ./

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g cmd/main.go

RUN go build -o main ./cmd/main.go

CMD ["./main"]
