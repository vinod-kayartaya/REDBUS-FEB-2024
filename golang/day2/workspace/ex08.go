package main

import (
	"fmt"
)

func main() {
	data := make([]int, 3, 5)

	fmt.Printf("len=%d, cap=%d, data=%v, first addr = %x\n", len(data), cap(data), data, &data[0])
	data = append(data, 10, 20, 1)
	fmt.Printf("len=%d, cap=%d, data=%v, first addr = %x\n", len(data), cap(data), data, &data[0])
	data = append(data, 100, 200, 29, 100, 200, 29, 100, 200, 29, 100, 200, 29)
	fmt.Printf("len=%d, cap=%d, data=%v, first addr = %x\n", len(data), cap(data), data, &data[0])
}
