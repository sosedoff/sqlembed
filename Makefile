.PHONY: build install test example

build:
	go build

test:
	go test ./...

install:
	go install

example:
	sqlembed -path=./example > example/queries.go
