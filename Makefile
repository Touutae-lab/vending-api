include .env

run-container:
	docker-compose --env-file .\.env -f .\docker\docker-compose.yml up -d

migrate-db:
	migrate -source file://docker/migrations -database "${POSTGRES_USER}://${POSTGRES_PASSWORD}:postgres@localhost:5432/${POSTGRES_DB}?sslmode=disable" up

migrate-db-down:
	migrate -source file://docker/migrations -database "${POSTGRES_USER}://${POSTGRES_PASSWORD}:postgres@localhost:5432/${POSTGRES_DB}?sslmode=disable" down

generate-jet-sql:
	 jet -dsn="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable" -schema=${DATABASE_SCHEMA_NAME} -path=./internal/database

generate-mock:
	mockgen -source .\internal\repository\product_repository.go -destination test/product_repository_mocks.go -package mocks
	mockgen -source .\internal\repository\machine_repository.go -destination test/machine_repository_mocks.go -package mocks
	mockgen -source .\internal\service\machine_service.go -destination test/machine_service_mocks.go -package mocks
	mockgen -source .\internal\service\product_service.go -destination test/product_service_mocks.go -package mocks
	mockgen -source .\internal\handler\product_handler_interface.go -destination test/product_handler_mocks.go -package mocks
	mockgen -source .\internal\handler\machine_handler_interface.go -destination test/machine_handler_mocks.go -package mocks
	mockgen -source .\internal\repository\user_repository.go -destination test/user_repository_mocks.go

generate-test-coverage:
	go test ./... -coverprofile="coverage.out"
	go tool cover -html="coverage.out" -o index.html

.PHONY: migrate-db run-container generate-jet-sql migrate-db-down generate-mock