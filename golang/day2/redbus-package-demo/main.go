package main

import (
	"fmt"

	"github.com/vinod-kayartaya/redbusutils/dateutils"
	"github.com/vinod-kayartaya/redbusutils/mathutils"
	"github.com/vinod-kayartaya/timeecho"
)

func main() {
	fmt.Println("Hello, world!")

	timeecho.EchoTime()
	y := -2023
	if leap, err := dateutils.IsLeap(y); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(y, "is a leap year -->", leap)
	}

	fmt.Println("Factorial of 5 is ", mathutils.Factorial(5))
}
