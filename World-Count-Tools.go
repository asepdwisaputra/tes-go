/*
	terdapat sebuah kalimat, kalimat dipecah perkata.
	apabila kata > 3 huruf dan memiliki minimal 1 fokal dihitung 1.
	tentukan total nilai dalam 1 kalimat.
*/

package main

import (
	"fmt"
	"strings"
)

func cekKata(char rune) bool {
	fokal := "aiueoAIUEO"
	return strings.ContainsRune(fokal, char)
}

func main() {
	s := "aku aku aku lagi makan ok ok suk$esya" // 1 kata bernilai 1
	total := 0
	//pecahan := strings.Fields(s) // OUTPUT sama
	pecahan := strings.Split(s, " ")

	for _, v := range pecahan {
		if len(v) >= 3 {
			Kata := false
			for _, w := range v {
				if cekKata(w) {
					Kata = true
					break
				}
			}
			if Kata {
				total++
			}
		}
	}
	fmt.Println(total)
}
