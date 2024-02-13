package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Example of handling errors")
	sum := 0.
	var nonNumericalInputs string

	for _, arg := range os.Args[1:] {
		// fmt.Printf("%T, %v\n", arg, arg)
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			sum += n
		} else {
			nonNumericalInputs += arg + ", "
		}
	}

	fmt.Printf("Sum of all input numbers is %v\n", sum)
	fmt.Println("You entered these non-numerical inputs -", nonNumericalInputs)
}
