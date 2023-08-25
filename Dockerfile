FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY internal ./internal
COPY pkg ./pkg
COPY config ./config
COPY docs ./docs
COPY cmd ./cmd

RUN go build -o /app/auth-service ./cmd/app

CMD ["/app/auth-service"]