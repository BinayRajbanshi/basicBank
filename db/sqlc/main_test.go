package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load the config")
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
