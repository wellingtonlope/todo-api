test:
	go test ./... -count=1

server:
	go run ./cmd/api/

build:
	go build ./cmd/api/
