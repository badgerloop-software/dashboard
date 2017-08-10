package main

import (
	"log"
	"net/http"
	"net"
	"fmt"
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

func initialize_UDP() {
	var ourAddrP *net.UDPAddr
	var err error
	ourAddrP, err = net.ResolveUDPAddr("udp", ":3000")
	CheckError(err)
	outAddr, err = net.ResolveUDPAddr("udp", "192.168.0.10:3000")
	CheckError(err)
	packet_conn, err = net.ListenUDP("udp", ourAddrP)
	CheckError(err)
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
				fmt.Println("got query packet")
				_, err = packet_conn.WriteToUDP([]byte("new phone who dis"), outAddr)
			}
			fmt.Print(string(buf[5:n]))
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
			fmt.Println("(Malformed packet, ", n, " bytes) ", buf[0:n], " from ", addr)
		}
		CheckError(err)
	}
}
/*****************************************************************************/
/*****************************************************************************/


/*****************************************************************************/
/*                                HTTP Handlers                              */
/*****************************************************************************/
func handler(w http.ResponseWriter, r *http.Request) {
	var err error
	testData := []models.Data{}
	err = database.GetConnection().Select(&testData, "SELECT * FROM Data LIMIT 1")
	CheckError(err)
	fmt.Fprintf(w, "(%s) %#v", r.URL.Path[1:], testData)
}

func UDPForwardingHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["data"]
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	/* valid API call */
	if message != nil {
		//fmt.Println("valid: ", message[0])
		w.WriteHeader(http.StatusOK)
		_, err := packet_conn.WriteToUDP([]byte(message[0]), outAddr)
		CheckError(err)
	/* invalid API call */
	} else {
		w.WriteHeader(http.StatusBadRequest)
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

	initialize_UDP()
	defer packet_conn.Close()

	/* Listen for microcontroller */
	go UDPServer()

	/* Serve on port 2000 */
	http.HandleFunc("/", handler)
	http.HandleFunc("/message", UDPForwardingHandler)
	log.Fatal(http.ListenAndServe(":2000", nil))
}

