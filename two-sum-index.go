// cetak 2 nilai index yang jika ditambahkan hsilnya = target

package main

import "fmt"

func twoSum(nums []int, target int) []int {

	for i := 0; i <= len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil // AWAS!!!
}
func main() {
	nums := []int{2, 7, 16, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
