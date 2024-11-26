all: build run

build:
	go build ./cmd/main

run:
	go run ./cmd/main

clean:
	go mod tidy

test_prepare:
	go get github.com/stretchr/testify/suite

test:
	go test -v -count=1 ./...