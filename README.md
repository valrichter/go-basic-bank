-GNU Make 4.3
-go version go1.21.0 linux/amd64
-golang-migrate 4.16.2
-Docker version 24.0.6, build ed223bc
-Docker image postgres 15.4
-sqlc v1.20.0

``` SELECT pg_terminate_backend(pg_stat_activity.pid)
FROM pg_stat_activity
WHERE pg_stat_activity.datname = 'bank'
AND pid <> pg_backend_pid(); ```

```SELECT usename, application_name, client_addr, backend_start, state
FROM pg_stat_activity
WHERE datname = 'bank';```

1. Creacion del DB schema
2. Configucion Postgres en Docker y creacion de la DB
3. Configucion golang-migrate para hacer migraciones de la DB
4. Configucion de sqlc y generacion de un CRUD basico para cada tabla con sqlc (SQL->[sqlc]->Go)