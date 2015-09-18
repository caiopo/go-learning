package main

import (
	"fmt"
	//	m "math"
	// "time"
)

func main() {

	c := make(chan int)

	go teste(10, c)

	go teste(5, c)

	go teste(8, c)

	fmt.Println(<-c, <-c, <-c)

}

func teste(i int, c chan int) {

	c <- i + 1

}
