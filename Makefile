.PHONY: all test server build format lint

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

