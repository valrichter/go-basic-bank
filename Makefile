postgres:
	docker run --name bank_db --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:15.4

createdb:
	docker exec -it postgres15.4 createdb --username=root --owner=root bank
	
dropdb:
	docker exec -it postgres15.4 dropdb bank

migrateAWS:
	migrate -path db/migration -database "postgresql://root:jZv5FR8Kl7zFLu0XeykL@basic-bank.c0yzhxhtjoab.sa-east-1.rds.amazonaws.com:5432/bank" -verbose up

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down 1
	
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/valrichter/go-basic-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock