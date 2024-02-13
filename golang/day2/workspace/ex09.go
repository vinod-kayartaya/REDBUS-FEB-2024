package main

import (
	"fmt"
)

func main() {
	nums := []int {1, 2, 3, 4, 5, 6, 7}

	fmt.Println(nums)
	nums = append(nums[:2],  nums[4:]...)
	fmt.Println(nums)
}
