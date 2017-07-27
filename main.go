package main

import (
	"log"
	//"net/http"
	"fmt"

	"github.com/badgerloop-software/dashboard/database"
	//api "github.com/badgerloop-software/dashboard/services"
	models "github.com/badgerloop-software/dashboard/models"
	//restful "github.com/emicklei/go-restful"
)

/*
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path accessed: %s", r.URL.Path[1:])
}
*/

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

	// Add the api data service
	//restful.Add(api.New())

	// Serve on port 2000
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":2000", nil))
}
