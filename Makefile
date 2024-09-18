build:
	go build -o bin/accounts ./cmd/main.go

run: build
	./bin/accounts $(ARGS)

.PHONY: build run
