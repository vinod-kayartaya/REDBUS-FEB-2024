package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", num, i, num*i)
	}
}
