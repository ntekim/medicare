ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack-cli.mk
include ./hack/hack.mk

DB_DSN=postgresql://nervix:pharmacist@localhost:5432/nervix_db?sslmode=disable


## gen/sqlc generate sql quries and related interfaces using sqlc
.PHONY: gen/sqlc
sqlc-gen:
	sqlc generate

# ==================================================================================== #
# SQL MIGRATIONS
# ==================================================================================== #

MIGRATION_PATH=./internal/dao/migrations
## migrations/new name=$1: create a new database migration
.PHONY: migrations/new
migrations-new:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -seq -ext=.sql -dir=${MIGRATION_PATH} ${name}

## migrations/up: apply all up database migrations
.PHONY: migrations/up
migrations-up:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=${MIGRATION_PATH} -database="${DB_DSN}" up

## migrations/down: apply all down database migrations
.PHONY: migrations/down
migrations/down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=${MIGRATION_PATH} -database="${DB_DSN}" down

## migrations/goto version=$1: migrate to a specific version number
.PHONY: migrations/goto
migrations/goto:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=${MIGRATION_PATH} -database="${DB_DSN}" goto ${version}

## migrations/force version=$1: force database migration
.PHONY: migrations/force
migrations/force:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=${MIGRATION_PATH} -database="${DB_DSN}" force ${version}

## migrations/version: print the current in-use migration version
.PHONY: migrations/version
migrations/version:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="${DB_DSN}" version

.PHONY: migrations/drop
migrations/drop:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=${MIGRATION_PATH} -database="${DB_DSN}" drop -f


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	make test
	go mod verify

## test run all the tests with covet
.PHONY: test
test:
	@GOFLAGS="-count=1" go test -v -cover -race -vet=off ./...