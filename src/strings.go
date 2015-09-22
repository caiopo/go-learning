package main

import (
	. "fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	t := time.Now().UTC().UnixNano()

	Println(t)

	Println(strconv.FormatInt(t, 10))
	str1, str2, str3 := "hist", "history", "tes te t asd fad h          "

	// array := strings.Fields(str3)
	// for _, s := range array {
	// 	Println(s)
	// }

	Println(strings.TrimPrefix(str2, str1), "\n", str3)

	Println(startsWith(str1, str2))

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
