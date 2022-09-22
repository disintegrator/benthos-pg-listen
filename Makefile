.PHONY: start-postgres
start-postgres:
	@./scripts/start-postgres.sh

.PHONY: stop-postgres
stop-postgres:
	@./scripts/stop-postgres.sh

.PHONY: psql
psql:
	@docker run --link listen-notify-demo -it --rm postgres:14 psql -h listen-notify-demo -U postgres -p 5432 -w
