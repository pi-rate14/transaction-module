package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/pi-rate14/transaction-module/api"
	db "github.com/pi-rate14/transaction-module/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/transaction_module?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
 
	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server")
	}
}