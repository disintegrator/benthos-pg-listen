.PHONY: start-postgres
start-postgres:
	@./scripts/start-postgres.sh

.PHONY: stop-postgres
stop-postgres:
	@./scripts/stop-postgres.sh

.PHONY: psql
psql:
	@./scripts/psql.sh
