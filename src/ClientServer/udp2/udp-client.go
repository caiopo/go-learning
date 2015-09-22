package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var targetPorts []string = []string{"56000", "56001", "56002"}

// const myPort = "55000"

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	fmt.Printf("Target ports: %s\n", targetPorts)

	for {

		var msg string

		// fmt.Print("Message to send: >> ")
		fmt.Scanf("%s", &msg)

		tempo := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

		for _, t := range targetPorts {
			go send(t, msg, tempo)
		}

		// time.Sleep(500 * time.Millisecond)

	}

}

func send(target, msg, tempo string) {

	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+target)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	// fmt.Println(port)

	defer Conn.Close()

	msg += ":" + tempo

	buf := []byte(msg)

	_, err = Conn.Write(buf)

	fmt.Printf("Sent message %s to %s\n", msg, target)
	if err != nil {
		fmt.Println(msg, err)
	}
}
