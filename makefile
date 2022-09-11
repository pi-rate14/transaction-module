# start the postgres container on port 5432
postgres:
	docker run --name postgres14simplebank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2

# create database
createdb:
	docker exec -it postgres14simplebank createdb --username=root --owner=root transaction_module

# drop database
dropdb:
	docker exec -it postgres14simplebank dropdb transaction_module

# perform up migration
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/transaction_module?sslmode=disable" -verbose up

# perform down migration
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/transaction_module?sslmode=disable" -verbose down

# sqlc 
sqlc:
	sqlc generate

# run unit tests
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb