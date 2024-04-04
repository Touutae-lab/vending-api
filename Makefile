include .env

run-container:
	docker-compose --env-file .\.env -f .\docker\docker-compose.yml up -d

migrate-db:
	migrate -source file://docker/migrations -database "${POSTGRES_USER}://${POSTGRES_PASSWORD}:postgres@localhost:5432/${POSTGRES_DB}?sslmode=disable" up

generate-jet-sql:
	 jet -dsn="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable" -schema=${DATABASE_SCHEMA_NAME} -path=./internal/database

generate-mock:
	mockgen -source .\internal\repository\product_repository.go -destination test/product_repository_mocks.go
	mockgen -source .\internal\repository\user_repository.go -destination test/user_repository_mocks.go
.PHONY: migrate-db run-container generate-jet-sql