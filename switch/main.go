package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("GO-SWITCH")

	// Basic

	option := 2
	fmt.Print("Write ", option, " as ")
	switch option {
	case 1:
		fmt.Println("Selected option 1")
	case 2:
		fmt.Println("Selected option 2")
	case 3:
		fmt.Println("selected option 3")
	}

	//  You can use commas to separate multiple expressions
	//  in the same case statement.
	// We use the optional default case in this example as well

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is an alternative way to express if/else logic
	// here we also show how the case expression can be non-contents

	currentTime := time.Now()
	switch {
	case currentTime.Hour() < 12:
		fmt.Println("Good Morning")
	case currentTime.Hour() > 12 && currentTime.Hour() < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good morning")
	}

	// switch expression used to types instead of values

	whatType := func(i interface{}) {
		switch reqType := i.(type) {
		case int:
			fmt.Println(" given type is integer")
		case bool:
			fmt.Println(" given type is boolean")
		default:
			fmt.Printf(" given type is unknown %T\n", reqType)
		}

	}

	whatType(false)
	whatType(12)
	whatType("hello!")
}
