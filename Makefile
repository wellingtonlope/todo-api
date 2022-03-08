test:
	go test ./...

server:
	go run ./cmd/http/server.go

build:
	go build ./cmd/http/server.go
