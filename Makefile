.PHONY: run
run:
	go run server.go

.PHONY: mod
mod:
	go mod download

.PHONY: test
test:
	go test ./...