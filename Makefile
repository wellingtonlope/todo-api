.PHONY: all test server build format lint swagger deps-update

all: format lint test

test:
	go test ./... -count=1

server:
	go run ./cmd/api/

build:
	go build ./cmd/api/

format:
	@if ! command -v gofumpt &> /dev/null; then \
		echo "gofumpt not found, installing..."; \
		go install mvdan.cc/gofumpt@latest; \
	fi
	gofumpt -w .

lint:
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run

swagger:
	@if ! command -v swag &> /dev/null; then \
		echo "swag not found, installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	swag init -g cmd/api/main.go -o docs/

deps-update:
	go get -u ./...
	go mod tidy

