package main

import (
	"database/sql"
	"fmt"
	"github.com/narymbaev/simple-bank/api"
	db "github.com/narymbaev/simple-bank/db/sqlc"
	"github.com/narymbaev/simple-bank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
	store := db.NewStore(conn)
	fmt.Printf("Address of x: %p\n", &store)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
