# Database

## Goose: setup database migrations

Run commands

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir=data/migrations create <migration-filename> sql
```

Edit generated migration file

```sql
-- +goose Up
-- +goose StatementBegin

<<< create table ... here >>>

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

<<< drop table ... here >>>

-- +goose StatementEnd
```

Run migration

```bash
goose -dir=data/migrations sqlite3 app.db up
# goose -dir=data/migrations sqlite3 app.db down
```

## SQLc: generate code from SQL schema

Edit [sqlc.yaml file](./../sqlc.yaml):

```yml
version: "2"
sql:
  - engine: "sqlite"
    queries: <queries-file-dir>
    schema: <goose-migration-folder-dir>
    gen:
      go:
        package: <file-output-location>
        out: <file-output-location>

```

Create queries file at the specified dir

```sql
-- name: GetBlog :one
SELECT * FROM blog 
WHERE id = ? LIMIT 1;

-- name: GetAllBlogs :many
SELECT * FROM blog 
ORDER BY time_created;

```

Run commands

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
sqlc generate
```