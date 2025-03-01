#!/bin/bash

set -a
# load goose environment variables
. .env
set +a

goose -dir data/migrations "${DB_DRIVER}" "${DB_CONNECTION_STRING}" "${@}"