

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

run:
	go run main.go

test: 
	go test ./...

build: 
	go build ./...

test-debug:
	go test ./... -v

test-cov:
	go test ./... -coverprofile=coverage.txt -covermode=atomic -race

multi-build:
	./scripts/multi-build.bash