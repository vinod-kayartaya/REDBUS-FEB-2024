package main

import (
	"fmt"
	"time"
)

func factorial(num int, ch chan int) {
	f := 1
	for num > 1 {
		f *= num
		num--
	}
	time.Sleep(3 * time.Second)
	ch <- f
}

func main() {
	ch := make(chan int)
	fmt.Println("creating a new goroutine")
	go factorial(5, ch)
	fmt.Println("new goroutine created")

	fmt.Println("waiting to receive from channel")
	fact := <-ch
	fmt.Println("got a value from channel")

	fmt.Printf("Factorial of 5 is %d\n", fact)
}
