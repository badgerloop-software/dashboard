package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Inits the connection to the database
// Should be called once from main routine
func InitDB(dataSourceName string) {
	var err error
	db, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
}

// Returns the pointer to the database connection
// Allows you to run queries
func GetConnection() *sqlx.DB {
	return db
}
