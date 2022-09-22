.PHONY: build
build:
	@go build -o benthos ./main.go

.PHONY: start-benthos
start-benthos: build
	@./benthos -c ./config.yml

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
