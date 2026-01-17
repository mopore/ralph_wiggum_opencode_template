.PHONY: test fmt build run

test:
	go test ./...

fmt:
	go fmt ./...

build:
	go build ./...

run:
	go run ./cmd/calculator
