.PHONY: test build run lint
test:
	go test ./... -race -cover
build:
	CGO_ENABLED=0 go build -o kvstore cmd/app/main.go
run:
	go run cmd/app/main.go
lint:
	golangci-lint run
