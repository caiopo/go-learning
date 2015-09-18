package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 0
	
	for i:=0; i < 10; i++{
		
		z = float64(z)-(math.Pow(float64(z), float64(2))-float64(x))/float64(2*x)

		fmt.Println(z)
		
	}

	return float64(z)
	
}

func main() {
	fmt.Println(Sqrt(2))
}
