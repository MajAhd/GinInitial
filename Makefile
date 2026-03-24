.PHONY: help dev build run test lint format tidy clean run-prod docker-build docker-prod

APP_NAME = server
BIN_DIR = bin

.DEFAULT_GOAL := help

## Show this help message
help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpDesc = substr(lastLine, 4); \
			printf "  \033[36m%-18s\033[0m %s\n", helpCommand, helpDesc; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

# Ensure go.sum exists
go.sum: go.mod
	go mod tidy

## Run development server with auto-reload using docker-compose
dev:
	docker compose up --build

## Build the production binary locally
build: go.sum clean
	mkdir -p $(BIN_DIR)
	go build -ldflags="-w -s" -o $(BIN_DIR)/$(APP_NAME) ./cmd/server

## Run the binary locally
run: build
	./$(BIN_DIR)/$(APP_NAME)

## Run the app in production-like mode directly
run-prod:
	GIN_MODE=release go run cmd/server/main.go

## Build production docker image
docker-build:
	docker build -t $(APP_NAME):latest .

## Run production docker container
docker-prod: docker-build
	docker run -p 8080:8080 -e GIN_MODE=release --env-file .env $(APP_NAME):latest

## Run tests
test: go.sum
	go test -v -cover ./...

## Run static analysis and linting
lint: go.sum
	go vet ./...
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed, running native vet only."; \
	fi

## Format code
format:
	go fmt ./...

## Tidy module reqs
tidy:
	go mod tidy

## Clean build artifacts
clean:
	rm -rf $(BIN_DIR)/ tmp/
