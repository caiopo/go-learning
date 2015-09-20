package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux! (\">)")

	case "darwin":
	default:
		fmt.Println("         ,")
		fmt.Println("Hipster (_)")

	default:
		fmt.Println(os)

	}
}
