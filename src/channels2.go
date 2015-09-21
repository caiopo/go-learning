package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go foo(ch)
	}

	fmt.Println("called foo")
	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {

		<-ch

		fmt.Println(i)
	}

	fmt.Println("gotcha")
}

func foo(ch chan bool) {

	fmt.Println("foo")

	time.Sleep(3 * time.Second)
	ch <- true
}
