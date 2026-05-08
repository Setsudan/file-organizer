APP_NAME=file-organizer

.PHONY: fmt vet test build run-once run-watch

fmt:
	gofmt -w .

vet:
	go vet ./...

test:
	go test -race -count=1 ./...

build:
	go build -o $(APP_NAME) ./cmd/file-organizer

run-once:
	go run ./cmd/file-organizer --mode=once

run-watch:
	go run ./cmd/file-organizer --mode=watch
