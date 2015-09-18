package main

import (
	"fmt"
	"time"
)

type Ponto struct {
	x int
	y int
}

func teste(){

	i := 10
	j := &i

	fmt.Println(i, j)

}


func main() {

	time.Sleep(5*time.Second)

	p := Ponto{1, 2}




	fmt.Println(p)

	v := &p

	p.x = 5

	fmt.Println(v)	
	
	
}