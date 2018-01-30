SHELL := /bin/sh

SRC := $(wildcard *.go)
RLSRC := graphql_collections.rl graphql_values.rl
GEN := $(RLSRC:.rl=.go)

.PHONY: all
all: lint.out vet.out coverage.out 

.PHONY: install
install: $(RLSRC) $(GEN) $(SRC)
	go install -v

cover.out: $(RLSRC) $(GEN) $(SRC)
	go test -v -cover -covermode atomic -coverprofile cover.out .

coverage.out: cover.out
	go tool cover -func=cover.out | tee coverage.out

vet.out: install
	go vet -v | tee vet.out

lint.out: install
	golint | tee lint.out

.PHONY: test
test: cover.out	

.PHONY: clean
clean:
	go clean -i ./...
	$(RM) *.out
	$(RM) $(GEN)

%.go: %.rl $(RLSRC)
	ragel -G2 -Z $<
