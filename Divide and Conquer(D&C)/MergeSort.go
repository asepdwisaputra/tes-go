package main

import "fmt"

func merge(left []int, right []int) []int {
	result := []int{}

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func main() {
	mySlice := []int{10, 5, 3, 7, 2, 8, 1, 6, 9, 4}

	sortedSlice := mergeSort(mySlice)

	fmt.Println(sortedSlice)
}
