/*
	urutkan data berikut:
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	sort.Ints(intSlice) // perintah urutkan
	fmt.Println("Sorted intSlice:", intSlice)

	strSlice := []string{"apple", "banana", "cherry", "date", "elderberry"}
	sort.Strings(strSlice) // perintah
	fmt.Println("Sorted strSlice:", strSlice)

}
