package main

import (
	"fmt"
	"time"
)

func main() {

	go teste()
	time.Sleep(10000)
	teste()

}

func teste() int {

	fmt.Println("oi")

	return 0

}
