package main

import (
    "fmt"
)

func main() {
	n1 := 1234
    fmt.Printf("type of n1 is %T, n1 = %v, address of n1 = %v\n", n1, n1, &n1)

	p1 := &n1
    fmt.Printf("type of p1 is %T, p1 = %v, value referred by p1 = %v\n", p1, p1, *p1)

	*p1 = 9000
    fmt.Printf("type of n1 is %T, n1 = %v, address of n1 = %v\n", n1, n1, &n1)

}

