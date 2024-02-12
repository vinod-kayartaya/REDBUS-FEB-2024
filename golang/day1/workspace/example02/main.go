package main

import "fmt"

func main() {

	myName := "Vinod"
	a, b := 100, 200
	var pi float64 = 3.14569
	// var tf bool
	tf := 123 > 29

	fmt.Printf("%T %s\n", myName, myName)
	fmt.Printf("%T %d\n", a, a)
	fmt.Printf("%T %d\n", b, b)
	fmt.Printf("%T %f\n", pi, pi)
	fmt.Printf("%T %v\n", tf, tf)
}