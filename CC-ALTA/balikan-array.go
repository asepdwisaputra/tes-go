package main

import "fmt"

func reverseArray(arr []int32) []int32 {
	// Write your code here

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func main() {
	arr := []int32{5, 3, 2, 4, 1} // 1,4,2,3,5
	fmt.Println(reverseArray(arr))
}
