package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	// Cara 1
	// prev := nums[0]
	// l := 1
	// for i := 1; i < len(nums); i++ {
	// 	if nums[i] != prev {
	// 		nums[l] = nums[i]
	// 		l++
	// 	}
	// 	prev = nums[i]
	// }
	// return l

	// Cara 2
	i := 0
	for j := range nums {
		if nums[i] != nums[j] {
			i += 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
