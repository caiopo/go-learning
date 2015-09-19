package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// const recvPort = ":56001"
// const targetPort = ":56002"

// const recvPort = ":56002"
// const targetPort = ":56001"

var (
	recvPort, targetPort, sendPort string
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

	hist = append(hist, "1", "2", "3")

	// fmt.Println(hist[0])

	// fmt.Println(has(hist, "5"))

	var sel int

	fmt.Scanf("%d", &sel)

	switch sel {
	case 0:
		recvPort = ":56000"
		targetPort = ":56001"
		sendPort = ":56020"
	case 1:
		recvPort = ":56001"
		targetPort = ":56000"
		sendPort = ":56021"

	}

	fmt.Printf("Listening on port %s\nSending to port %s\nMy sending port is %s\n", recvPort, targetPort, sendPort)

	ch := make(chan string, 3)

	go recv(ch)

	for {
		select {
		case msg := <-ch:
			if has(hist, msg) {
				continue
			} else {
				hist = append(hist, msg)
				go send(msg)
			}
			// fmt.Printf(">%s<\n", msg)

			// default:
			// 	continue
		}

	}

}

func recv(ch chan string) {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", recvPort)
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

func send(msg string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+targetPort)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+sendPort)
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	// fmt.Println(port)
	defer fmt.Println("Finished sending", msg)
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
