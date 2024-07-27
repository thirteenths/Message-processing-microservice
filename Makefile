lint:
	golangci-lint run -v -c golangci.yml ./...

migration-create-sql:
	goose -dir=./migrations create $(NAME) sql

swag-generate-api:
	swag init --parseDependency --parseInternal  -g internal/ports/https/message.go -o cmd/api/docs/ -g cmd/api/main.go

run-api:
	go run cmd/api/main.go