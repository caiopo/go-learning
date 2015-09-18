package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {

	fmt.Println("Launching server...")
	for {
		// connected := true
		connect()

		time.Sleep(5 * time.Second)
	}
}

func connect() int {

	fmt.Println("Listening port 8081")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	handleConnection(conn)

	return 0
	// run loop forever (or until ctrl-c)

} // - See more at: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/#sthash.WBTeQDyu.dpuf

func handleConnection(conn net.Conn) int {

	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			// connected = false
			return 0
		}
		// output message received
		fmt.Println("Message Received:", string(message), "Error:", err, "\n")

		// sample process for string received
		newmessage := strings.ToUpper(message)

		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}

}
