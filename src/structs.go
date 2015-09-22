package main

import (
	. "fmt"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	key   string
	milis int64
}

func main() {
	str := "teste " + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	Println(str)

	m := make(map[string]int64)

	a := strings.Fields(str)

	Println(a)

	Println(a[1])

	// x, y := strconv.Atoi(a[1])

	// Println("xy:", x, y)

	value, _ := strconv.ParseInt(a[1], 10, 64)

	m[a[0]] = value

	Println(m)

}
