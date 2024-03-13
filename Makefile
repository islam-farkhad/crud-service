include .env

ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=$(DB_USER) password=$(PASSWORD) dbname=$(DBNAME) host=localhost port=$(PORT) sslmode=$(SSLMODE)
endif

MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: start-test-service
start-test-service:
	docker-compose up -d

.PHONY: build-app-image
build-app-image:
	docker build -t posts-service:1.0 ./

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: run-integration-tests
run-integration-tests:
	go test -v -tags=integration ./tests

.PHONY: run-unit-tests
run-unit-tests:
	go test -v ./internal/handlers/...
