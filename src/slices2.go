package main

import (
	. "fmt"
	"time"
)

func main() {

	a := append([]int64{time.Now().UTC().UnixNano()}, []int64{1, 2, 3, 4}...)

	Println(a)

}
