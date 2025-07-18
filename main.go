package main

import (
	"context"
	"log"

	"github.com/BinayRajbanshi/GoBasicBank/api"
	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	serverAddress := config.ServerAddress
	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
