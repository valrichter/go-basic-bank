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

``` SELECT usename, application_name, client_addr, backend_start, state
FROM pg_stat_activity
WHERE datname = 'bank'; ```

1. Creacion del DB schema
2. Configucion Postgres en Docker y creacion de la DB
3. Configucion golang-migrate para hacer migraciones de la DB
4. Configucion de sqlc y generacion de un CRUD basico para cada tabla con sqlc (SQL->[sqlc]->Go)
5. Generacion de datos ficticios y creacion de test unitarios para los CRUD de las tablas account, entrie, transfer
6. Creacion de store_procedure para la transferencia de dinero entre usuarios y su respectivo test
7. Creacion de test con go routines para tracciones concurrentes y evitar transaction lock junto con el manejo de deadlocks
8. Creacion de test de deadlocks y modificacion del codigo para evitar deadlocks
9. Creacion de Continouous Integration (CI) con GitHub Actions para automatizacion de test 