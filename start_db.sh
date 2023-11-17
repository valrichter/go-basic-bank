#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
DB_SOURCE=postgresql://root:secret@postgres:5432/bank?sslmode=disable
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up


echo "start the app"

exec "$@"