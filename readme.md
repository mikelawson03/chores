Things to install:
Goose (go install github.com/pressly/goose/v3/cmd/goose@latest)
Sqlite3 (go get github.com/mattn/go-sqlite3)

Document:
how migrations work
how to run them manually
where the migration files live
BUT the server will create the DB and run the migrations automatically on start


Endpoints


## Scheduling

The scheduler automatically generates daily, weekly, and monthly chore assignments within a configurable planning horizon.

See `docs/scheduler.md` for detailed scheduling behavior and cadence rules.