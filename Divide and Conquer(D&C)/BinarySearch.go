package main

import "fmt"

func binarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target1 := 5
	target2 := 11

	fmt.Printf("Hasil pencarian %d: %v\n", target1, binarySearch(mySlice, target1)) // Output: true (elemen 5 ditemukan)
	fmt.Printf("Hasil pencarian %d: %v\n", target2, binarySearch(mySlice, target2)) // Output: false (elemen 11 tidak ditemukan)
}
