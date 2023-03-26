package main

import "fmt"

func main() {
	fmt.Println("For loop...")

	// WAY: 1
	fmt.Println("WAY 1 SIMPLE LOOP")

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// WAY: 2 --> like while loop
	fmt.Println("WAY 2 LIKE WHILE LOOP")
	j := 11

	for j < 20 {
		fmt.Println(j)
		j++
	}

	// WAY: 3 --> like Infinite for loop
	fmt.Println("WAY 3 LIKE Infinite for loop")
	k := 20
	for {
		fmt.Println(k)
		if k == 30 {
			break
		}
		k++
	}

	// WAY: 4 --> to iterate over elements in a collection,
	fmt.Println("WAY 3 Range")
	list := []int{31, 32, 33, 34, 35, 36, 37, 38, 39, 40}

	for index, value := range list {
		fmt.Println(index, value)
	}

}
