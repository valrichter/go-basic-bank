# Proyecto de Backend extenso con Go<img id="go" src="https://devicon-website.vercel.app/api/go/plain.svg?color=%2300ACD7" width="30" />

## Herramietas usadas:
- GNU Make 4.3
- go version go1.21.0 linux/amd64
- golang-migrate 4.16.2
- Docker version 24.0.6
- Docker image postgres:15.4
- sqlc v1.20.0

## Acciones realizadas en el proyecto:
1. Creacion del DB schema
   <iframe width="560" height="315" src='https://dbdiagram.io/embed/64f10dee02bd1c4a5ec5a954'> </iframe>
2. Configuracion de PostgreSQL con Docker y creacion de la DB
3. Configuracion golang-migrate para hacer migraciones de la DB
4. Configuracion de sqlc para hacer consultas SQL con codigo Go. Generacion de un CRUD basico para las tablas account, entry & transfer con sqlc. Como funciona: SQL->[sqlc]->Go
5. Generacion de datos falsos y creacion de test unitarios para los CRUD de las tablas account, entrie, transfer
6. Creacion de store_procedure para la transferencia de dinero entre usuarios y su respectivo test
7. Creacion de test con go routines para tracciones concurrentes y evitar transaction lock junto con el manejo de deadlocks
8. Creacion de test de deadlocks y modificacion del codigo para evitar deadlocks
9.  Creacion de Continouous Integration (CI) con GitHub Actions para mayor automatizacion

## Comandos para desarrollo
Desconcetar todos los usarios de la DB:  
``` SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'bank' AND pid <> pg_backend_pid(); ```

Cuantos usarios esta conectados a la DB:  
``` SELECT usename, application_name, client_addr, backend_start, state FROM pg_stat_activity WHERE datname = 'bank'; ```