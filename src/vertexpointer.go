package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Deslocar(val float64) {

	v.X += val
	v.Y += val
}

func main() {
	ponto := Vertex{2, 3}

	fmt.Println(ponto)

	ponto.Deslocar(3)

	fmt.Println(ponto)

}
