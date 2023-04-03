package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	oneDArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("1d:", oneDArray)

	var twoDArray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDArray[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDArray)
	fmt.Println("2d array len: ", len(twoDArray))
}
