package main

import "fmt"

func main() {
	array := [...]int{1, 3, 4, 2, 7, 5, 8, 6}

	var temp int

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				// Tukar nilai array
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}

	fmt.Println(array)
}
