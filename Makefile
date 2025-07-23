postgres:
	docker run --name postgres17 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root basic_bank

dropdb:
	docker exec -it postgres17 dropdb basic_bank

migrateup:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5433/basic_bank?sslmode=disable" -verbose up

migratedown: 
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5433/basic_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/BinayRajbanshi/GoBasicBank/db/sqlc Store

#In a Makefile, the .PHONY declaration is used to mark targets as phony, meaning they are not actual files, but rather commands or actions that should always be executed when called.
.PHONY: createdb postgres migrateup migratedown dropdb sqlc test server 

 