postgres:
	docker run --name postgres15.4 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:15.4

createdb:
	docker exec -it postgres15.4 createdb --username=root --owner=root bank
	
dropdb:
	docker exec -it postgres15.4 dropdb bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down
.PHONY: postgres createdb dropdb