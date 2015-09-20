package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// const recvPort = ":56001"
// const targetPorts = ":56002"

// const recvPort = ":56002"
// const targetPorts = ":56001"

var (
	serverName, recvPort string
	targetPorts          []string
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

var hist []string

func main() {

	hist = make([]string, 0)

	// fmt.Println(hist[0])

	// fmt.Println(has(hist, "5"))

	var sel int

	fmt.Scanf("%d", &sel)

	switch sel {
	case 0:
		recvPort = "56000"
		targetPorts = []string{"56001", "56002"}
		serverName = "Server 0"
	case 1:
		recvPort = "56001"
		targetPorts = []string{"56000", "56002"}
		serverName = "Server 1"
	case 2:
		recvPort = "56002"
		targetPorts = []string{"56001", "56000"}
		serverName = "Server 2"
	}

	fmt.Printf("My name is %s\nListening on port %s\nSending to ports %s\n", serverName, recvPort, targetPorts)

	go showHist()

	ch := make(chan string, 3)

	go recv(ch)

	for {
		select {
		case msg := <-ch:
			if has(hist, msg) {
				continue
			} else {
				hist = append(hist, msg)
				for _, t := range targetPorts {
					go send(t, msg)
				}
			}
			// fmt.Printf(">%s<\n", msg)

			// default:
			// 	continue
		}

	}

}

func recv(ch chan string) {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+recvPort)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)

		msgRecv := string(buf[0:n])
		fmt.Println("Received ", msgRecv, " from ", addr)
		ch <- msgRecv
		time.Sleep(1 * time.Second)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}

func send(target, msg string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+target)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	// fmt.Println(port)
	defer fmt.Printf("Finished sending %s to %s\n", msg, target)
	defer Conn.Close()

	// for {

	buf := []byte(msg)
	_, err = Conn.Write(buf)
	if err != nil {
		fmt.Println(msg, err)
	}

	// time.Sleep(time.Second * 1)
	// }
}

func has(h []string, str string) bool {

	for _, s := range h {
		if s == str {
			return true
		}
	}

	return false

}

func showHist() {

	var in string

	for {

		fmt.Scanf("%s", &in)

		// if in != nil {

		fmt.Println(hist)

		// }
	}
}

// func send(msg string, ch chan string) {
// 	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:56001")
// 	CheckError(err)

// 	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
// 	CheckError(err)

// 	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
// 	CheckError(err)

// 	// fmt.Println(port)

// 	defer Conn.Close()
// 	for {

// 		msg := <-ch

// 		buf := []byte(msg)
// 		_, err := Conn.Write(buf)
// 		if err != nil {
// 			fmt.Println(msg, err)
// 		}
// 		time.Sleep(time.Second * 1)
// 	}
// }
