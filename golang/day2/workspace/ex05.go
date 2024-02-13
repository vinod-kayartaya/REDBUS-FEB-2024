package main

import (
	"fmt"
)

func doSomething() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var n, d int

	fmt.Printf("Enter numerator: ")
	fmt.Scan(&n)

	fmt.Printf("Enter denominator: ")
	fmt.Scan(&d)

	q := n / d
	fmt.Printf("%v / %v is %v\n", n, d, q)

}
func main() {
	fmt.Println("Start of main()")
	doSomething()
	fmt.Println("End of main()")
}
