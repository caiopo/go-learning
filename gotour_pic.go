package main

//to use in https://tour.golang.org/moretypes/14

import (
	"fmt"
	// m "math"
	// "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	// matrix := make([][]uint8, 0, dy)

	// for i := 0; i < dy; i++ {

	// 	for j := 0; j < dx; j++ {

	// 		matrix[j][i] = uint8(j) ^ uint8(i)

	// 	}
	// }

	// return matrix

	matrix := make([][]uint8, 0, dy)

	for i := 0; i < dy; i++ {

		line := make([]uint8, 0, dx)

		for j := 0; j < dx; j++ {

			// line[j] = m.Pow(i, j)

			line = append(line, uint8((i+1)*(j+1)))

		}

		matrix = append(matrix, line)

	}

	return matrix

}

func main() {
	// pic.Show(Pic)

	dx, dy := 10, 10

	m := Pic(dx, dy)

	for i := 0; i < dy; i++ {
		fmt.Println(m[i])
	}

}
