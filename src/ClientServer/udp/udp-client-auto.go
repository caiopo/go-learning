package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const filename = "test.txt"

var targetPorts []string = []string{"56000", "56001", "56002"}

// const myPort = "55000"

func getFile() []string {
	file, err := ioutil.ReadFile(filename)
	CheckError(err)

	s := strings.Fields(string(file))

	return s
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Printf("Target ports: %s\n", targetPorts)

	file := getFile()

	fmt.Println(file)

	tamArray := len(file)

	var msg string

	for {

		msg = file[rand.Intn(tamArray)] + file[rand.Intn(tamArray)]

		tempo := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

		fmt.Printf("Sent message: %s\n", msg)

		for _, t := range targetPorts {
			go send(t, msg, tempo)
		}

		time.Sleep(time.Millisecond)

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

	// fmt.Printf("Sent message %s to %s\n", msg, target)
	if err != nil {
		fmt.Println(msg, err)
	}
}
