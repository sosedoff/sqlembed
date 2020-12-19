.PHONY: install example

install:
	go install

example:
	sqlembed -path=./example > example/queries.go
