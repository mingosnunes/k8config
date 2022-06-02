

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

run:
	go run main.go

test: 
	go test ./...

test-debug:
	go test ./... -v

test-cov:
	go test ./... -cover