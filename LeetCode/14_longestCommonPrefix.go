package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	// if len(strs) == 0 {
	// 	return ""
	// }
	// if len(strs) == 1 {
	// 	return strs[0]
	// }

	// // Temukan panjang prefix minimum dari semua string
	// minLen := len(strs[0])
	// for _, str := range strs {
	// 	if len(str) < minLen {
	// 		minLen = len(str)
	// 	}
	// }

	// // Loop melalui setiap indeks sampai panjang minimum
	// for i := 0; i < minLen; i++ {
	// 	for j := 1; j < len(strs); j++ {
	// 		if strs[j][i] != strs[0][i] {
	// 			return strs[0][:i]
	// 		}
	// 	}
	// }

	// return strs[0][:minLen]

	////////////////

	// Cara Simpel
	pref := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], pref) {
			pref = pref[:len(pref)-1]
			fmt.Println(pref)
		}
		//fmt.Println(strs[i])
	}
	return pref
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
