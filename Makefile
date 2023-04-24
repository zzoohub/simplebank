postgres:
	docker run --name postgres15 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1033 -d postgres:15-alpine

createdb:  
	docker exec -it postgres15 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres15 dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/simplebank?sslmode=disable" up

migrateup1:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/simplebank?sslmode=disable" up 1


migratedown:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/simplebank?sslmode=disable" down

migratedown1:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/simplebank?sslmode=disable" down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/storg.go github.com/zzoopro/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 test server mock