package main

import (
	"fmt"
)

func main() {

	// a := make([]int, 5, 5)

	// fmt.Println(a)

	// b := make([]int, 3, 3)

	// a = append(a, 1, 2, 3, 4)

	// a = append(a, b...)

	// fmt.Println(a)

	// c := []int{1, 2, 3}

	// fmt.Println(c)

	s := []string{"abc", "def"}

	var msg string

	for _, i := range s {

		msg += i + " "

	}

	d := []byte(msg)

	fmt.Println(s, "\n", d)
}
