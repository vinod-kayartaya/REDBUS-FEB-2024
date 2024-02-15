package main

import (
	"fmt"
)

var ch chan Factorial

type Factorial struct {
	n int
	f int
}

func findFactorial(num int) {
	f := 1
	for i := 1; i <= num; i++ {
		f *= i
	}
	ch <- Factorial{num, f}
}

func main() {
	ch = make(chan Factorial, 50)
	nums := []int{10, 5, 8, 1, 4, 12, 2, 12}

	for _, n := range nums {
		go findFactorial(n)
	}

	for i := 1; i < len(nums); i++ {
		obj := <-ch
		fmt.Printf("factorial of %d is %d\n", obj.n, obj.f)
	}
}
