package main

import "fmt"

func main() {

	var a, b, c int

	a = 1
	b = 1

	var sa = string(a)
	var sb = string(b)

	fmt.Printf(sa + "\n" + sb)

	for i := 0; i < 15; i++ {

		c = a + b

		a = b

		b = c

		fmt.Println(string(c))

	}

}
