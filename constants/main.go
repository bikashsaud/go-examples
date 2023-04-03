package main

import (
	"fmt"
	"math"
)

const version float64 = 1.20

func main() {
	fmt.Println("GO CONSTANTS")

	fmt.Printf("Go version: %f", version)

	const PI = 3.14
	fmt.Printf("\nPI: %f", PI)

	const n = 3e7

	const d = 3e20 / n
	fmt.Println("\n ", d)

	fmt.Println("\n ", int64(d))

	// define multiple constants
	const (
		Age       = 30
		Name      = "John"
		IsMarried = false
	)

	fmt.Println("\n ", Age)
	fmt.Println("\n ", Name)
	fmt.Println("\n ", IsMarried)

	// sin
	fmt.Println(math.Sin(90))
}
