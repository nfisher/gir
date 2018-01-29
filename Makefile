SHELL := /bin/sh

.PHONY: all
all:
	ragel -G2 -Z graphql_collections.rl 
	go install -v
	go vet -v
	go test -v

.PHONY: clean
clean:
	go clean -i ./...
