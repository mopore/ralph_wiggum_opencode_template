.PHONY: test fmt run build

test:
	go test ./...

fmt:
	gofmt -w .

run:
	go run ./cmd/calculator

build:
	go build -o bin/calculator ./cmd/calculator
