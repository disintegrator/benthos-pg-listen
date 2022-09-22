.PHONY: build
build: main.go
	@go build -o benthos ./main.go

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
