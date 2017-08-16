package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

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

/*****************************************************************************/
/*                     Microcontroller-related Networking                    */
/*****************************************************************************/
var outAddr *net.UDPAddr
var packet_conn *net.UDPConn
var mcuBuffer *bytes.Buffer

func initialize_UDP() {
	var ourAddrP *net.UDPAddr
	var err error

	ourAddrP, err = net.ResolveUDPAddr("udp", ":3000")
	CheckError(err)
	outAddr, err = net.ResolveUDPAddr("udp", "192.168.0.10:3000")
	CheckError(err)
	packet_conn, err = net.ListenUDP("udp", ourAddrP)
	CheckError(err)

	/* initialize empty buffer */
	mcuBuffer = bytes.NewBuffer(make([]byte, 0))
}

func UDPServer() {

	var err error
	var n int
	dat := models.Data{}
	buf := make([]byte, 1024)
	var addr *net.UDPAddr

	fmt.Println("Starting UDP Server")

	for {
		n, addr, err = packet_conn.ReadFromUDP(buf[:])

		/* Message, not a packet */
		if n > 5 && buf[0] == 'M' && buf[1] == 'S' && buf[2] == 'G' {
			/* respond to microcontroller querying for dashboard */
			if strings.Contains(string(buf[0:n]), "dashboard?") {
				_, err = packet_conn.WriteToUDP([]byte("new phone who dis"), outAddr)
			} else {
				fmt.Print(string(buf[5:n]))
				mcuBuffer.Write(buf[5:n])
			}
			/* SpaceX Packet */
		} else if n == 34 {
			dat, err = models.ParseSpaceXPacket(buf[:34])
			if err == nil {
				models.PrintSpaceX(dat)
			}
			/* Dashboard Packet */
		} else if n == 47 {
			dat, err = models.ParseDashboardPacket(buf[:47])
			if err == nil {
				models.PrintDashboard(dat)
				// TODO: push to DB
			}
			/* Malformed Packet*/
		} else {
			fmt.Println("(Malformed packet, ", n, " bytes) ", string(buf[0:n]), " from ", addr)
		}
		CheckError(err)
	}
}

/*****************************************************************************/
/*****************************************************************************/

/*****************************************************************************/
/*                                HTTP Handlers                              */
/*****************************************************************************/
/*
 * Database Querying API
 * Purpose: view X latest data points from the database that contains
 *          microcontroller sensor data
 * Usage: localhost:2000
 */
func handler(w http.ResponseWriter, r *http.Request) {

	var err error
	testData := []models.Data{}

	// TODO: how many data points do we want?
	err = database.GetConnection().Select(&testData, "SELECT * FROM Dashboard.Data ORDER BY created DESC LIMIT 1;")
	CheckError(err)
	myTestData, err := json.Marshal(testData)
	CheckError(err)
	// why would we need to Marshal and then Unmarshal?
	err = json.Unmarshal(myTestData, &testData)
	CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(myTestData)
}

/*
 * Microcontroller Command-Forwarding API
 * Purpose: run commands on microcontroller remotely
 * Usage: localhost:2000/message?data=<command>
 */
func UDPForwardingHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["data"]
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	/* microcontroller-command API call */
	if message != nil {
		fmt.Println("valid: ", message[0])
		w.WriteHeader(http.StatusOK)
		_, err := packet_conn.WriteToUDP([]byte(message[0]), outAddr)
		CheckError(err)
		/* invalid API call */
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

/*
 * Microcontroller Output Buffer API
 * Purpose: display console outputs from microcontroller in browser
 * Usage: localhost:2000/data[?reset]
 */
func bufferRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	/* reset buffer */
	if r.URL.Query()["reset"] != nil {
		mcuBuffer.Reset()
		w.WriteHeader(http.StatusOK)
		/* send buffer contents */
	} else {
		w.Write(mcuBuffer.Bytes())
	}
}

/*****************************************************************************/
/*****************************************************************************/

func db_test() {
	var err error
	testData := []models.Data{}
	err = database.GetConnection().Select(&testData, "SELECT * FROM Data")
	CheckError(err)
	fmt.Printf("query returned %d results.\n", len(testData))
}

func main() {

	/* Setup database connection */
	database.InitDB("dashboard:betsy@tcp(badgerloop.com:3306)/Dashboard")
	db_test()

	db, err := sql.Open("mysql", "dashboard:betsy@tcp(badgerloop.com:3306)/Dashboard")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO Data VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? , ?)")
	if err != nil {
		panic(err.Error())
	}

	defer stmtIns.Close()

	_, err = stmtIns.Exec(101, "yoyo", 14, 67, 10745, 9592, 295, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true}, sql.NullInt64{Int64: 42, Valid: true})
	if err != nil {
		panic(err.Error())
	}

	initialize_UDP()
	defer packet_conn.Close()

	/* Listen for microcontroller */
	go UDPServer()

	/* Serve on port 2000 */
	http.HandleFunc("/", handler)
	http.HandleFunc("/message", UDPForwardingHandler)
	http.HandleFunc("/buffer", bufferRequestHandler)
	log.Fatal(http.ListenAndServe(":2000", nil))
}
