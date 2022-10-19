.PHONY: build
build:
	@go build -o bin/basic ./cmd/basic/main.go
	@go build -o bin/benthos ./cmd/benthos/main.go

.PHONY: start-basic
start-basic: build
	@./bin/basic -dsn postgres://postgres@localhost:5432/postgres -channel outbox_items__insert

.PHONY: start-benthos
start-benthos: build
	@./bin/benthos --env-file .env -c ./config.yml

.PHONY: start-mock
start-mock:
	@npm i
	@npm start

.PHONY: start-postgres
start-postgres:
	@./scripts/start-postgres.sh

.PHONY: stop-postgres
stop-postgres:
	@./scripts/stop-postgres.sh

.PHONY: psql
psql:
	@./scripts/psql.sh

.PHONY: insert-item
insert-item:
	@./scripts/insert-item.sh
