# PostgreSQL LISTEN input for Benthos

This project demonstrates how to build a Benthos plugin to work with PostgreSQL's `LISTEN` feature. For comparison, a barebones Go program is included which implements the same functionality.

## Prerequisites

- Docker
- Go 1.19+

## Build

- Clone this repository
- Start a toy PostgreSQL server: `make start-postgres`
- Start the custom Benthos build: `make start-benthos`
- _(Optional)_ start the basic program for comparison: `make start-basic`
- In a separate terminal, insert fake outbox items: `make insert-item`
  - Alternatively, jump into a `psql` session and work with postgres interactively: `make psql`
- When you're done, terminate the PostgreSQL server: `make stop-postgres`

## Notable source files

- `cmd/basic/main.go`: The entrypoints for the basic Go program
- `cmd/benthos/main.go`: The entrypoints for the custom Benthos program
- `internal/postgres/input_listen.go`: The custom Benthos plugin to implement a PostgreSQL LISTEN input
- `postgres/docker-entrypoint-initdb.d/create_outbox.sql`: A script that is used to bootstrap a new PostgreSQL server
