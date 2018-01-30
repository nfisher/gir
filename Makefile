SHELL := /bin/sh

SRC := $(wildcard *.go)

.PHONY: all
all: lint.out vet.out coverage.out 

.PHONY: install
install: graphql_collections.go $(SRC)
	go install -v

cover.out: $(SRC)
	go test -v -cover -covermode atomic -coverprofile cover.out .

coverage.out: cover.out
	go tool cover -func=cover.out | tee coverage.out

vet.out: install
	go vet -v | tee vet.out

lint.out: install
	golint | tee lint.out

.PHONY: clean
clean:
	go clean -i ./...

graphql_collections.go: graphql_collections.rl
	ragel -G2 -Z graphql_collections.rl
