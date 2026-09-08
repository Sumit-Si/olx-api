.PHONY: build run

build:
# 	@go build -o bin/main main.go
	@go build -o bin/api ./cmd/api

run: build
# 	@./bin/main
	@./bin/api