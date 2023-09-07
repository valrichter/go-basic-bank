-GNU Make 4.3
-go version go1.21.0 linux/amd64
-golang-migrate 4.16.2
-Docker version 24.0.6, build ed223bc
-Docker image postgres 15.4

``` SELECT pg_terminate_backend(pg_stat_activity.pid)
FROM pg_stat_activity
WHERE pg_stat_activity.datname = 'bank'
AND pid <> pg_backend_pid(); ```

```SELECT usename, application_name, client_addr, backend_start, state
FROM pg_stat_activity
WHERE datname = 'bank';```