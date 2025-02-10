.PHONY: all build run test generate

all: build

build:
	GOOS=linux GOARCH=amd64 go build -o shortener cmd/main.go

run:
	docker-compose up --build

generate:
	protoc --go_out=. --go-grpc_out=. proto/shorturl.proto

test:
	go test ./tests/...
