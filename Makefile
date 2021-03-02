#

PG_URL?='postgres://postgres:postgres@localhost/customers?sslmode=disable'

init-db:
	@echo "==> init database schema"
	@psql $(PG_URL) -v ON_ERROR_STOP=1 -f sql/init/init_db.sql
.PHONY: init-db

migrate:
	@echo "=> Migrate"
	@go run ./cmd/migrations/migration.go -dir=./sql/migrations postgres ${PG_URL} up
.PHONY: migrate
