package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// ConnectionCredentials struct that defines the database connection credentials from environment variables
type ConnectionCredentials struct {
	USERDB     string
	PASSWORDDB string
	HOSTDB     string
	PORTDB     string
	DBNAME     string
}

//ValidateCredentials method
//Method that verifies that the credentials are not empty
func (cc ConnectionCredentials) ValidateCredentials() {
	if cc.USERDB == "" {
		fmt.Println("[Rap API][Error] The environment variable USERDB is undefined.")
		os.Exit(1)
	}
	if cc.PASSWORDDB == "" {
		fmt.Println("[Rap API][Error] The environment variable PASSWORDDB is undefined.")
		os.Exit(1)
	}
	if cc.HOSTDB == "" {
		fmt.Println("[Rap API][Error] The environment variable HOSTDB is undefined.")
		os.Exit(1)
	}
	if cc.PORTDB == "" {
		fmt.Println("[Rap API][Error] The environment variable PORTDB is undefined.")
		os.Exit(1)
	}
	if cc.DBNAME == "" {
		fmt.Println("[Rap API][Error] The environment variable DBNAME is undefined.")
		os.Exit(1)
	}
}

//GetConnection function
//Function that allows to coneect to a postgres database
func GetConnection() *sql.DB {
	cc := ConnectionCredentials{
		USERDB:     "admin_rap",
		PASSWORDDB: "septimojinete",
		HOSTDB:     "127.0.0.1",
		PORTDB:     "5432",
		DBNAME:     "db_rap",
	}

	cc.ValidateCredentials()

	ccString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cc.USERDB,
		cc.PASSWORDDB,
		cc.HOSTDB,
		cc.PORTDB,
		cc.DBNAME,
	)

	db, err := sql.Open("postgres", ccString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
