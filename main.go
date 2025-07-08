package main

import (
	"context"
	"log"

	"github.com/BinayRajbanshi/GoBasicBank/api"
	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	serverAddress := "0.0.0.0:8080"
	connPool, err := pgxpool.New(context.Background(), "postgresql://root:secret@localhost:5433/basic_bank?sslmode=disable")

	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
