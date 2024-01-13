DB_URL=postgresql://root:secret@localhost:5432/bank?sslmode=disable

postgres:
	docker run --name bank_db --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:15.4-alpine

db_create:
	docker exec -it bank_db createdb --username=root --owner=root bank
	
db_drop:
	docker exec -it bank_db dropdb bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup_1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown_1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml 

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/valrichter/go-basic-bank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/valrichter/go-basic-bank/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name go-basic-bank-redis -p 6379:6379 -d redis:7.2.3-alpine

.PHONY: postgres db_create db_drop migrateup migratedown migrateup_1 migratedown_1 new_migration db_docs db_schema sqlc test server mock proto evans redis