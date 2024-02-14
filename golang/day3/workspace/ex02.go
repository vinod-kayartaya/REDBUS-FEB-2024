package main

import (
	"fmt"
)

func increment(n *int) {
	*n += 10
	fmt.Printf("value and addr of n in `increment()` is %v and %v\n", *n, n)
}

func main() {
	num := 1000
	fmt.Printf("num = %v, address of num is %v\n", num, &num)
	increment(&num)
	fmt.Printf("num = %v, address of num is %v\n", num, &num)
	increment(&num)
	fmt.Printf("num = %v, address of num is %v\n", num, &num)

}
