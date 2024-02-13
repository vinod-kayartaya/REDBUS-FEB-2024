package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example usage of an array")
	// var nums [5]int

	nums := [5]int{10, 20, 30, 40, 50} // cannot be done with `var` declration
	fmt.Println("nums is", nums)

	fmt.Println("Enter 5 numbers: ")
	for i := 0; i < len(nums); i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println("nums is", nums)

	for index, value := range nums {
		fmt.Printf("nums[%d] = %d\n", index, value)
	}
}
