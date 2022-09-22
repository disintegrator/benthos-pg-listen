# PostgreSQL LISTEN input for Benthos

This project demonstrates how to build a Benthos plugin to work with
PostgreSQL's `LISTEN` feature.

## Build

- Clone this repository
- Start a toy PostgreSQL server with `make start-postgres`
- Start the custom Benthos build `make start-benthos`
- In a separate terminal, insert fake outbox items `make insert-item`
  - Alternatively, jump into a `psql` session and work with postgres interactively: `make psql`
