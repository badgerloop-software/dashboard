package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/badgerloop-software/dashboard/database"
	api "github.com/badgerloop-software/dashboard/services"
	models "github.com/badgerloop-software/dashboard/models"
	restful "github.com/emicklei/go-restful"
)

func handler(w http.ResponseWriter, r *http.Request) {

	var err error
	testData := []models.Data{}

	err = database.GetConnection().Select(&testData, "SELECT * FROM Data LIMIT 1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "(%s) %#v", r.URL.Path[1:], testData)
}

func main() {

	var err error

	// Setup database connection
	database.InitDB("dashboard:betsy@tcp(badgerloop.com:3306)/Dashboard")

	testData := []models.Data{}

	err = database.GetConnection().Select(&testData, "SELECT * FROM Data")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("query returned %d results.\n", len(testData))
	fmt.Printf("Example: %#v\n", testData[0])

	// Add the api data service
	restful.Add(api.New())

	// Serve on port 2000
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":2000", nil))
}
