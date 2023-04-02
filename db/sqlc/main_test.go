package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	deSorce = "postgresql://root:1033@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	
	testDB, err = sql.Open(dbDriver, deSorce)
	if err != nil {
		log.Fatal("cannot connect to database.")
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}