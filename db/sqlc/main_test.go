package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), "postgresql://root:secret@localhost:5433/basic_bank?sslmode=disable")

	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
