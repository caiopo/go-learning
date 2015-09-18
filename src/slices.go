package main

import (
	"fmt"
)

func main() {

	a := make([]int, 5, 5)

	fmt.Println(a)

	b := make([]int, 3, 3)

	a = append(a, 1, 2, 3, 4)

	a = append(a, b...)

	fmt.Println(a)

}
