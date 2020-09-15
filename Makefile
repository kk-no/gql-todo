# Go

.PHONY: run
run: 
	go run cmd/gql-todo/main.go

.PHONY: test
test: 
	go test ./...

.PHONY: tidy
tidy: 
	go mod tidy

.PHONY: mod
mod: 
	go mod download

# GQL

.PHONY: gql-gen
gql-gen:
	cd graphql; \
	go run github.com/99designs/gqlgen generate