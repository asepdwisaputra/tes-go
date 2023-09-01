package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 { // handle only 1 element
		return strs[0]
	}

	// sort them first, the most different one will be in first and last
	sort.Strings(strs)
	//fmt.Println(strs)

	// compare first and last
	l := len(strs)
	////////////////////////////

	/*
		1 | 0 | f | f
		2 | 1 | l | l
		3 | 2 | o | o
		4 | 3 | w | r
		````
		flo
	*/
	for i := range strs[0] {
		if strs[0][i] != strs[l-1][i] {
			return strs[0][:i]
			// DIBACA!!!
			// mencetal index[0] dari index awal sampai index(i-1)
			// karena mencetak[literasi mulai : literasi akhir-1]
		}
	}

	/////////////////////////////
	return strs[0]
}

func main() {
	strs := []string{"flow", "flor"}
	fmt.Println(longestCommonPrefix(strs)) //  flo => kesamaan

	// strs2 := []string{"flow", "flor", "flare", "flan", "flix", "floor"}
	// fmt.Println(longestCommonPrefix(strs2)) // fl

	// strs3 := []string{"flow", "flor", "fix", "fidio", "fan", "fans", "fuyoo"}
	// fmt.Println(longestCommonPrefix(strs3)) //f
}
