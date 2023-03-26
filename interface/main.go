package main

import "fmt"

func main() {
	var any interface{}

	any = "Hello world!"

	fmt.Println(any)
	// output: hello world!

	any = 42
	fmt.Println(any)
	// output: 42!

	any = true
	println(any)

	interfaceFunc("Hello Earth!")
	interfaceFunc(43)
	interfaceFunc(false)
	arr := [...]string{"Bikash", "bishal", "Binash"}
	interfaceFunc(arr)

}

func interfaceFunc(v interface{}) {
	fmt.Println(v)
}

