package main

import "fmt"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Examples.")

	var user User
	user.Name = "John Wick"
	user.Age = 29

	fmt.Printf("User name: %s user age: %d", user.Name, user.Age)

	fmt.Println("&user", &user)

	fmt.Println("*user", *&user)

	fmt.Println("User info")

	var u *User = &user

	fmt.Println("u= ", u)
	fmt.Println("*u= ", *u)

}
