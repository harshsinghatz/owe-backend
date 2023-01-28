package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:harsh@localhost:5432/owe?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M){
	conn, err:= sql.Open(dbDriver,dbSource);

	if err!=nil {
		log.Fatal("Something went wrong while connecting to database")
	}

	testQueries = New(conn)

	// m.Run returns failed test cases so if test cases are passed it would return with status 0 which means success
	os.Exit(m.Run())
}