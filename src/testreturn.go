package main

import "fmt"

func main() {
	b := foo()

	fmt.Println(b(), b(), b())

}

func foo() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
