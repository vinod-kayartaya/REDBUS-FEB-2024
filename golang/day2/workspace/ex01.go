package main

import (
	"errors"
	"fmt"
)

func factorial(num int) (int, error) {
	if num < 0 {
		// panic("This func cannot handle negative inputs")
		return 0, errors.New("`factorial()` cannot handle negative inputs")
	}
	f := 1
	for num > 1 {
		f *= num
		num--
	}
	return f, nil
}

func main() {
	fmt.Println("Error handling demo")
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if fact, err := factorial(n); err != nil {
		fmt.Println("OOPS! there was an error: ", err.Error())
	} else {
		fmt.Printf("Factorial of %v is %v\n", n, fact)
	}

	fmt.Println("End of main()")
}
