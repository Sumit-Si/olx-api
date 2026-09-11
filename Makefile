.PHONY: build run

build:
# 	@go build -o bin/main main.go
	@go build -o bin/api ./cmd/api

run: build
# 	@./bin/main
	@./bin/api

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down