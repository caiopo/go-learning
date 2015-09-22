package main

import (
	. "fmt"
	"reflect"
)

func main() {

	a := 42

	Println(a, reflect.TypeOf(a))

	b := 42.0

	Println(b, reflect.TypeOf(b))

	c := "42"

	Println(c, reflect.TypeOf(c))

}
