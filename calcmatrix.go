package main

import (
	"fmt"
	// "math"
)

type Matrix [][]int

// func (m *Matrix) det() {
func det(m Matrix) {

	for _, i := range m {

		for _, j := range i {

			// m[i][j] = i + j

			fmt.Println(i, j)
		}
	}

}

func main() {

	// a = make([][]int)

	a := make(Matrix, 0, 5)

	// a[0] = {1, 2, 3, 4, 5}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = i + j
		}
	}

	fmt.Println(a)

	// a.det()
	det(a)

	fmt.Println(a)
}
