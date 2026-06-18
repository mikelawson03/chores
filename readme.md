## Getting Started
### Install:
Goose 
```go install github.com/pressly/goose/v3/cmd/goose@latest```

Sqlite3 
```go get github.com/mattn/go-sqlite3```

## Set up
Create a `.env` file in the project root, and include two variables:

DB_PATH - relative path to DB
PORT - port for server to use

Example:
```
DB_PATH="./data/chores.db"
PORT=8080
```

## Database and Migrations
The server uses a Sqlite database for all data persistence. The server will create the DB and run the migrations automatically on startup. If you receive an error that the DB cannot be created, ensure you have created a folder at the path indicated by your DB_PATH env variable.

Migrations can be found in `sql/schema`. If, for any reason, you need to run up or down migrations, you can do so by using the following commands in the project root:

### Up migration:
```
goose sqlite3 -dir ./sql/schema {DB_PATH} up
```
Note: an up migration will run ALL migrations in the folder up to current. To run one at a time, you can specific `up-by-one`

### Down migration:
```
goose sqlite3 -dir ./sql/schema {DB_PATH} down
```
Note: a down migration will roll back a single migration (e.g., v3 to v2). To roll back to the beginning, use `reset` rather than `down`

## Endpoints

For full documentation of endpoints and API feature, see `docs/api.md`

## Scheduling

The scheduler automatically generates daily, weekly, and monthly chore assignments within a configurable planning horizon.

See `docs/scheduler.md` for detailed scheduling behavior and cadence rules.

## Balancing

The balancer assigns chore assignments to available users, attempting to keep workloads balanced across the planning window.

See `docs/balancer.md` for detailed balancing behavior and assignment rules.

## Status: Server MVP Complete

Implemented:
- Chore templates
- Assignment management
- Automated scheduling
- Workload balancing
- SQLite persistence
- REST API

Planned:
- Client applications
- Automated scheduler execution
- Notifications
- Calendar integration
- Household management modules