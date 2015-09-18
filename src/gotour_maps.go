package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	// a = make([]string, 0, )

	a := strings.Fields(s)

	for i := 0; i < len(a); i++ {

		m[a[i]] = m[a[i]] + 1

	}

	return m
	// map[string]int{"x": 1}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(">> ")

	input, _ := reader.ReadString('\n')

	fmt.Println(WordCount(input))
}
