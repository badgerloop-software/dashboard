package main

import (
	"log"
	"net/http"

	"github.com/badgerloop-software/dashboard/database"
	api "github.com/badgerloop-software/dashboard/services"
	restful "github.com/emicklei/go-restful"
)

func main() {
	// Setup database connection
	database.InitDB("dashboard:betsy@tcp(96.42.32.19:3306)/Dashboard")

	// Add the api data service
	restful.Add(api.New())

	// Serve on port 2000
	log.Fatal(http.ListenAndServe(":2000", nil))
}
