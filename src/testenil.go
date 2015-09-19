package main

import (
	"fmt"
	"time"
)

func main() {

	go teste(1)
	time.Sleep(1 * time.Millisecond)
	teste(2)

}

func teste(num int) int {

	fmt.Println("oi", num)

	return 0

}
