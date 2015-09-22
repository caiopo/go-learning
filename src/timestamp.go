package main

import (
	. "fmt"
	"time"
)

func main() {
	for {

		Println(time.Now().UTC().UnixNano())
	}
}
