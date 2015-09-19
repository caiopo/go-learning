package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// const recvPort = ":56001"
// const sendPort = ":56002"

const recvPort = ":56002"
const sendPort = ":56001"

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	ch := make(chan string)

	go recv(ch)

	for {
		select {
		case msg := <-ch:

			go send(msg, ch)

			fmt.Printf(">%s<\n", msg)

		default:
			continue
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

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}

func send(msg string, ch chan string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+sendPort)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	// fmt.Println(port)

	defer Conn.Close()
	for {

		msg := <-ch

		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
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
