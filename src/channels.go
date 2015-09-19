package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go foo(ch)

	time.Sleep(time.Millisecond * 50)

	select {
	case msg := <-ch:
		fmt.Printf(">%s<\n", msg)
	default:
		fmt.Println("Empty!")

	}

	// fmt.Printf(">%s<", msg)

}

func foo(ch chan string) {
	ch <- "teste"
}
