package main

import (
	"fmt"
	//	m "math"
	"time"
)

type Ponto struct {
	x, y int
}

func main() {

	fmt.Println(time.Now())

	c := make(chan Ponto)

	p1 := Ponto{1, 1}

	go teste(p1, c)

	fmt.Println(<-c)

	// fmt.Println(i)

}

func teste(v Ponto, c chan Ponto) {

	v.x += 1
	v.y += 1

	c <- v

}
