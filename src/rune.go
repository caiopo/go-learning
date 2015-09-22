package main

import (
	. "fmt"
	// "reflect"
	"strings"
)

func filterColon(x rune) bool {
	return Sprintf("%c", x) == ":"
}

// func runeToString(r rune) string {
// 	return
// }

func main() {

	// Println(byte(s[3]), reflect.TypeOf(byte(s[3])))

	// Println([]byte(":")[0])

	// Sprintf(format, ...)

	// Println(reflect.TypeOf(":"))

	s := "a:b:c"

	Println(&s)

	Println(strings.FieldsFunc(s, filterColon))

}
