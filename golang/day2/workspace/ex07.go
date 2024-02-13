package main

import (
	"fmt"
)

func main() {
	fmt.Println("")

	var arr [10]int
	x := arr[3:5]

	fmt.Printf("%d %d\n", len(x), cap(x))

	nums := []int{10, 20, 30}
	fmt.Printf("len(nums) = %v\n", len(nums))
	fmt.Printf("cap(nums) = %v\n", cap(nums))

	fmt.Printf("Address of nums is %#x\n", &nums[0])
	fmt.Println(nums)
	nums = append(nums, 12)
	fmt.Printf("Address of nums is %#x\n", &nums[0])
	fmt.Println(nums)
}
