package main

import (
	. "fmt"
)

func main() {
	m := make(map[string]int)

	m["teste"] = 42

	// Println(m)

	for i, j := range m {
		Println(i, j)
	}

}
