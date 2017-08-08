package main

import (
	"log"
	"net/http"
	"net"
	"fmt"

	"github.com/badgerloop-software/dashboard/database"
	//api "github.com/badgerloop-software/dashboard/services"
	models "github.com/badgerloop-software/dashboard/models"
	//restful "github.com/emicklei/go-restful"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	var err error
	testData := []models.Data{}

	err = database.GetConnection().Select(&testData, "SELECT * FROM Data LIMIT 1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "(%s) %#v", r.URL.Path[1:], testData)
}

func UDPServer() {

	dat := models.Data{}

	addr, err := net.ResolveUDPAddr("udp", ":3000")
	CheckError(err)

	outAddr, err := net.ResolveUDPAddr("udp", "192.168.0.10:3000")
	CheckError(err)

	conn, err := net.ListenUDP("udp", addr)
	CheckError(err)

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf[:])
		/* SpaceX Packet */
		if n == 34 {
			dat, err = models.ParseSpaceXPacket(buf[:34])
			if err == nil {
				models.PrintSpaceX(dat)
			}

		/* Dashboard Packet */
		} else if n == 47 {
			dat, err = models.ParseDashboardPacket(buf[:47])
			if err == nil {
				models.PrintDashboard(dat)
			}
			// TODO: push to DB

		/* Malformed Packet*/
		} else {
			fmt.Println("(Malformed packet, ", n, " bytes) ", buf[0:n], " from ", addr)
		}
		CheckError(err)

		n, err = conn.WriteToUDP([]byte("I got your packet (:"), outAddr)
		CheckError(err)
	}
}

func main() {

	var err error

	// Setup database connection
	database.InitDB("dashboard:betsy@tcp(badgerloop.com:3306)/Dashboard")

	testData := []models.Data{}

	err = database.GetConnection().Select(&testData, "SELECT * FROM Data")
	CheckError(err)

	fmt.Printf("query returned %d results.\n", len(testData))
	//fmt.Printf("Example: %#v\n", testData[0])

	// Add the api data service
	//restful.Add(api.New())

	// Serve on port 2000
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":2000", nil))

	UDPServer()
}
