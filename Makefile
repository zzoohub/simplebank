postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1033 -d postgres:15-alpine

createdb:  
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/simple_bank?sslmode=disable" up

migratedown:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/simple_bank?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown test