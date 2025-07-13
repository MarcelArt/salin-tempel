go:
	@go run main.go

dev:
	@air

swag:
	@swag init --parseDependency --parseInternal

build-win:
	@GOOS=windows GOARCH=amd64 go build -o bin/papan-klip.exe main.go

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/papan-klip main.go

build:
	@make build-linux
	@make build-win