package main

import (
	"fmt"
	"sort"
)

func main() {
	// Inisialisasi array string
	angka := []int{-1, 2, 1, 5, 3, 6, 4}
	names := []string{"John", "Mary", "Peter", "Jane"}
	names2 := []string{"John", "Mary", "Peter", "Jane"}

	// Urutkan array
	sort.Ints(angka)

	sort.Slice(names, func(a, b int) bool {
		return names[a] < names[b]
	})

	sort.Strings(names2)

	// Cetak array yang telah diurutkan
	fmt.Println(angka)
	fmt.Println(names)
	fmt.Println(names2)

}
