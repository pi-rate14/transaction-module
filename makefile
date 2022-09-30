# start the postgres container on port 5432
postgres:
	docker run --name postgres14simplebank --network transaction-module-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2

# create database
createdb:
	docker exec -it postgres14simplebank createdb --username=root --owner=root transaction_module

# drop database
dropdb:
	docker exec -it postgres14simplebank dropdb transaction_module

# perform up migration
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/transaction_module?sslmode=disable" -verbose up

# perform last up migration
migrateupone:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/transaction_module?sslmode=disable" -verbose up 1

# perform down migration
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/transaction_module?sslmode=disable" -verbose down

# perform last down migration
migratedownone:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/transaction_module?sslmode=disable" -verbose down 1

# sqlc 
sqlc:
	sqlc generate

# run unit tests
test:
	go test -v -cover ./...

# start HTTP server
server:
	go run main.go

# mock data
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/pi-rate14/transaction-module/db/sqlc Store

# run docker container on network
# sudo docker run --name transaction-module --network transaction-module-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres14simplebank:5432/transaction_module?sslmode=disable" transaction-module

.PHONY: postgres createdb dropdb migrateup migrateupone migratedown migratedownone sqlc test server mock