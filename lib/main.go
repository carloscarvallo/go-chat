package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const port = "8080"

// RunHost takes an ip as argument
// and listens for connections on that ip
func RunHost(ip string) {

	ipAndPort := ip + ":" + port

	listener, listenErr := net.Listen("tcp", ipAndPort)

	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}

	conn, acceptErr := listener.Accept()

	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received: ", message)
}

// RunGuest takes destination ip
// as argument and connects to that ip
func RunGuest(ip string) {

}
