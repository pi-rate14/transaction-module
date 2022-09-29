package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/pi-rate14/transaction-module/api"
	db "github.com/pi-rate14/transaction-module/db/sqlc"
	"github.com/pi-rate14/transaction-module/util"
)



func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load env variables")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config,store)
	if err != nil {
		log.Fatal("cannot create server")
	}
 
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server")
	}
}