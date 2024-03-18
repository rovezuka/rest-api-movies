FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go build -o app

FROM alpine

WORKDIR /app

COPY . .

CMD ["./app"]