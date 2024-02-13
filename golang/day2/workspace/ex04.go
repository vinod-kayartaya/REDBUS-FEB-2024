package main

import (
	"fmt"
)

func main() {
	greet := func() {
		fmt.Println("Hello, world!")
	}
	fmt.Printf("type of greet is %T\n", greet)
	greet()
	welcome := greet
	welcome()

	func() {
		fmt.Println("This is an example anonymous function")
	}()
}
