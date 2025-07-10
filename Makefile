go:
	@go run main.go

dev:
	@air

swag:
	@swag init --parseDependency --parseInternal

build:
	@go build -o bin/papan-klip main.go