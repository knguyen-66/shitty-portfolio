# Data & migrations

This directory contains the data for the page.

Blog posts are stored in the [blogs](./blogs/) directory, in Markdown format with some custom formatting.

The [migrations](./migrations/) directory contains the SQL migrations for the database schema, including the initial database schema.

The generation of database code is handled by [sqlc](https://github.com/sqlc-dev/sqlc) and [goose](https://github.com/pressly/goose) library.

You can edit the sqlc configuration in the [sqlc.yaml](../sqlc.yaml) file.

## For updating the database schema

- Create new migrations with goose:

    ```bash
    cd data/migrations 
    goose create WHATEVER_MIGRATION_CONTEXT_FILENAME sql
    ```

- Edit the created migration file in the `data/migrations` directory.
- Run the migration with goose to apply the changes to the database:

    ```bash
    cd data/migrations 
    goose up
    # or
    goose up-to MIGRATION_FILENAME
    ```

## For reverting the database migrations

- Run the migration with goose to revert the changes to the database:

    ```bash
    cd data/migrations
    goose down
    # or
    goose down-to MIGRATION_FILENAME
    ```

## For updating the database queries

- Edit/add/remove SQL queries in the [query.sql](./query.sql) file.
- Run the sqlc command:

    ```bash
    sqlc generate
    ```
