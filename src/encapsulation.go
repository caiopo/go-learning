package main

import "fmt"

var teste string = "oi"

func main() {
	teste := "bar"
	foo()
	fmt.Println("in main", teste)
}

func foo() {

	fmt.Println("in foo", teste)

}
