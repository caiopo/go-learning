package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	key   string
	milis int64
}

// const recvPort = ":56001"
// const targetPorts = ":56002"

// const recvPort = ":56002"
// const targetPorts = ":56001"

var (
	serverName, recvPortS, recvPortC string
	targetPorts                      []string
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

var hist map[string]int64

func main() {

	hist = make(map[string]int64, 0)

	sel := os.Args[1]

	switch sel {
	case "3":
		recvPortC = "56000"
		recvPortS = "56100"
		targetPorts = []string{"56101", "56102"}
		serverName = "Server 0"
	case "1":
		recvPortC = "56001"
		recvPortS = "56101"
		targetPorts = []string{"56100", "56102"}
		serverName = "Server 1"
	case "2":
		recvPortC = "56002"
		recvPortS = "56102"
		targetPorts = []string{"56101", "56100"}
		serverName = "Server 2"
	}

	fmt.Printf("My name is %s\nListening clients on port %s\nListening servers on port %s\nSending to ports %s\n", serverName, recvPortC, recvPortS, targetPorts)

	go showHist()

	ch := make(chan Entry, 3)

	go recvFromClient(ch)

	go recvFromServer()

	go sendAndSleep()

	updateHist(ch)

}

func updateHist(ch <-chan Entry) {
	for {
		select {
		case msg := <-ch:

			// fmt.Println("Update history:", msg)

			if hist[msg.key] != 0 {
				if hist[msg.key] < msg.milis {

					hist[msg.key] = msg.milis

				}
			} else {
				hist[msg.key] = msg.milis
			}

		}

	}
}

func recvFromClient(ch chan<- Entry) {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+recvPortC)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		// recebe um array de bytes
		n, _, err := ServerConn.ReadFromUDP(buf)

		//converte o array de bytes para string
		msgRecv := string(buf[0:n])
		// fmt.Println("Received ", msgRecv, " from ", addr)

		// separa a mensagem e o timestamp
		if len(msgRecv) > 3 {
			tempRecv := strings.FieldsFunc(msgRecv, filterColon)

			// cria uma entry com os dados recebidos e a coloca no canal
			parsed, _ := strconv.ParseInt(tempRecv[1], 10, 64)

			entryRecv := Entry{tempRecv[0], parsed}

			ch <- entryRecv
		}
		// time.Sleep(1 * time.Second)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}

func recvFromServer() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+recvPortS)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		// recebe um array de bytes
		n, _, err := ServerConn.ReadFromUDP(buf)

		//converte o array de bytes para string
		msgRecv := string(buf[0:n])
		// fmt.Println("Received ", msgRecv, " from ", addr)

		// separa a mensagem e o timestamp

		// time.Sleep(1 * time.Second)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		syncHist(msgRecv)
	}

}

func syncHist(msg string) {

	if startsWith("history ", msg) {

		strings.TrimPrefix(msg, "history")

		array := strings.Fields(msg)

		m := make(map[string]int64)

		if len(array) >= 2 {
			for _, i := range array {

				aTemp := strings.FieldsFunc(i, filterColon)

				if len(aTemp) >= 2 {
					m[aTemp[0]], _ = strconv.ParseInt(aTemp[1], 10, 64)
				}
			}

			for k, v := range m {
				if hist[k] == 0 || hist[k] < v {

					hist[k] = v

				}
			}
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

func contains(h []string, str string) bool {

	for _, s := range h {
		if s == str {
			return true
		}
	}

	return false

}

func sendAndSleep() {
	for {

		// time.Sleep(500 * time.Millisecond)
		time.Sleep(time.Second * 2)
		for _, t := range targetPorts {
			sendHistory(t)
		}

	}
}

func sendHistory(target string) {

	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+target)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()

	msg := "history "

	for k, v := range hist {

		msg += k + ":" + strconv.FormatInt(v, 10) + " "

	}

	// fmt.Println("Sending history:", msg)

	buf := []byte(msg)
	_, err = Conn.Write(buf)

	// fmt.Printf("Sent history to %s\n", target)

	if err != nil {
		fmt.Println(msg, err)
	}

}

func startsWith(smaller, larger string) bool {

	lenS := len(smaller)
	lenL := len(larger)

	switch {

	case lenS > lenL:
		return false

	case smaller == larger[0:lenS]:
		return true

	default:
		return false
	}

}

func filterColon(x rune) bool {
	return fmt.Sprintf("%c", x) == ":"
}

type EntryList []Entry

func (e EntryList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e EntryList) Len() int {
	return len(e)
}

func (e EntryList) Less(i, j int) bool {
	return e[i].milis < e[j].milis
}

func sortMapByValue(m map[string]int64) EntryList {
	e := make(EntryList, len(m))
	i := 0
	for k, v := range m {
		e[i] = Entry{k, v}
		i++
	}

	// sort.Sort(e)
	sorted := true

	for sorted {
		sorted = false
		for i := 0; i < e.Len()-1; i++ {

			if e.Less(i+1, i) {
				sorted = true
				e.Swap(i, i+1)

			}
		}
	}

	return e
}

func showHist() {

	var in string

	for {

		fmt.Scanf("%s", &in)

		sortedArray := sortMapByValue(hist)

		fmt.Print("History:\n{ ")
		for _, i := range sortedArray {
			fmt.Printf("%s ", i.key)
		}
		fmt.Println("}")
	}
}

// func makeHist(m map[string]int64) {

// }

// func sortMapByValue(m map[string]int64) EntryList {
// 	e := make(EntryList, len(m))
// 	i := 0
// 	for k, v := range m {
// 		e[i] = Entry{k, v}
// 	}

// 	sort.Sort(e)
// 	return e
// }

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

// func updateHist(ch <-chan string) {
// 	for {
// 		select {
// 		case msg := <-ch:

// 			if startsWith("history ", msg) {

// 				tempHist := strings.Fields(strings.TrimPrefix(msg, "history "))

// 				// if len(tempHist) >= len(hist) {
// 				if hist[0] < tempHist[0] {
// 					hist = tempHist

// 					for _, t := range targetPorts {
// 						go sendHistory(t)
// 					}

// 				}

// 			} else {

// 				if contains(hist, msg) {
// 					continue
// 				} else {
// 					hist = append(hist, msg)
// 					for _, t := range targetPorts {
// 						go sendHistory(t)
// 					}
// 				}
// 			}
// 			// fmt.Printf(">%s<\n", msg)

// 			// default:
// 			// 	continue
// 		}

// 	}
// }
