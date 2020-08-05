package pgsql12

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "Pa$$@word1"
	dbname   = "testgo"
)

var Db *sql.DB

// Init initial database connection
func Init() error {
	var err error

	pgSQLConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", pgSQLConn)
	// defer db.Close()

	if err != nil {
		log.Printf("sql.Open() %v", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("db.Ping() %v", err)
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	log.Println("Connected to database")

	Db = db

	return nil

}
