package main

import (
	"log"
	"net/http"

	"github.com/badgerloop-software/dashboard/database"
)

func main() {
	// Setup database connection
	database.InitDB("dashboard:betsy@tcp(96.42.32.19:3306)/Dashboard")

	// Stores connection to database, can run queries, more of placeholder
	db := database.GetConnection()

	// Serve on port 2000
	log.Fatal(http.ListenAndServe(":2000", nil))

	db.Close()
}
