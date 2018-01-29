.PHONY: all
all:
	ragel -G2 -Z graphql_collections.rl 
	go test -v
