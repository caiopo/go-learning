package main

import (
	"fmt"
	"net"
	// "strconv"
	"time"
)

const targetPort = "56001"
const myPort = "55000"

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	fmt.Printf("My port: %s\nTarget port: %s\n", myPort, targetPort)

	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+targetPort)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+myPort)
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	// fmt.Println(port)

	defer Conn.Close()
	for {

		var msg string

		fmt.Scanf("%s", &msg)

		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
