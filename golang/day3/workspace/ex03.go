package main

import (
	"fmt"
)

func getNames() *[]string {
	names := []string{"Vinod", "Shyam", "John", "Jane"}
	return &names
}

func doSomething() {
	someNames := getNames()
	fmt.Println("Names are", *someNames)
	fmt.Printf("Address of someNames is %v\n", &(*someNames)[0])
	// someNames is a local variable, and will be removed at this time
	// and the allocated heap memory is going to be garbage collected
	// (since this is the only pointer variable to that location)
}

func main() {
	doSomething()
	doSomething()
	doSomething()
}
