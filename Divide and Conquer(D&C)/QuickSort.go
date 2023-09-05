package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main() {
	mySlice := []int{10, 5, 3, 7, 2, 8, 1, 6, 9, 4}

	sortedSlice := quickSort(mySlice)

	fmt.Println(sortedSlice)
}
