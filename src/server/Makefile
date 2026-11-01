# Optional: load variables (e.g. DB_ADDR) from .env when it exists.
-include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: run
run: ## Run the API (use `air` instead for live reload)
	@go run ./cmd/api

.PHONY: build
build: ## Build the API binary into ./bin
	@go build -o ./bin/api ./cmd/api

.PHONY: test
test: ## Run all tests
	@go test -v ./...

.PHONY: lint
lint: ## Run go vet
	@go vet ./...

.PHONY: migration
migration: ## Create a new migration: make migration <name>
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up: ## Apply all pending migrations
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down: ## Roll back migrations: make migrate-down <n>
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: gen-docs
gen-docs: ## Regenerate Swagger docs in ./docs
	@swag init -g ./api/main.go -d cmd,internal && swag fmt

.PHONY: docker-build
docker-build: ## Build the API Docker image
	@docker build -t weather-api .

# Swallow extra goals used as arguments (e.g. `make migration add_x`).
%:
	@:
