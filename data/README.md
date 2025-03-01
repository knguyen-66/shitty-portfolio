# Data & migrations

This directory contains the data for the page.

Blog posts are stored in the [blogs](./blogs/) directory, in Markdown format with some custom formatting.

The [migrations](./migrations/) directory contains the SQL migrations for the database schema, including the initial database schema.

The generation of database code is handled by [SQLc](https://github.com/sqlc-dev/sqlc) and [goose](https://github.com/pressly/goose) library.

You can edit the SQLc configuration in the [sqlc.yaml](../sqlc.yaml) file.

## For updating the database schema

- Create new migrations with goose:

    ```bash
    ./.scripts/goose.sh create WHATEVER_MIGRATION_CONTEXT_FILENAME sql
    ```

- Edit the created migration file in the `data/migrations` directory. The annotation list are defined in the [goose documentation](https://pressly.github.io/goose/documentation/annotations/).
- Run the goose migration command to apply the changes to the database:

    ```bash
    ./.scripts/goose.sh up
    # or
    ./.scripts/goose.sh up-to MIGRATION_FILENAME
    ```

## For reverting the database migrations

- Run the goose migration command to revert the changes to the database:

    ```bash
    ./.scripts/goose.sh down
    # or
    ./.scripts/goose.sh down-to MIGRATION_FILENAME
    ```

## For updating the database queries

- Edit/add/remove SQL queries in the [query.sql](./query.sql) file.

    ```sql
    -- name: QUERY_FUNCTION_NAME :QUERY_RESULT_TYPE(one/many/exec/execresult/...)
    QUERY_STRING

    -- name: ANOTHER_QUERY_FUNCTION_NAME :QUERY_RESULT_TYPE
    QUERY_STRING_2

    ...
    ```

- Run the SQLc command to update the generated code:

    ```bash
    sqlc generate
    ```
