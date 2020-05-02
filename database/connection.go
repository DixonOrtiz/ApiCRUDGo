package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// ConnectionCredentials struct that defines the database connection credentials from environment variables
type ConnectionCredentials struct {
	USERDB     string
	PASSWORDDB string
	HOSTDB     string
	PORTDB     string
	DBNAME     string
}

//GetConnection function
//Function that allows to coneect to a postgres database
func GetConnection() *sql.DB {
	cc := ConnectionCredentials{
		USERDB:     os.Getenv("USERDB"),
		PASSWORDDB: os.Getenv("PASSWORDDB"),
		HOSTDB:     os.Getenv("HOSTDB"),
		PORTDB:     os.Getenv("PORTDB"),
		DBNAME:     os.Getenv("DBNAME"),
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cc.USERDB,
		cc.PASSWORDDB,
		cc.HOSTDB,
		cc.PORTDB,
		cc.DBNAME,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
