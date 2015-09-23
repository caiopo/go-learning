package main

import (
	. "fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// file, err := os.Open("readfiles.txt")
	file, err := ioutil.ReadFile("readfiles.txt")
	check(err)

	s := string(file)
	Println(strings.Fields(s))

}
