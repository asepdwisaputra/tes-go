package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'countValidWords' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func countValidWords(s string) int32 {
	// Write your code here
	total := 0
	kataperKata := strings.Split(s, " ")
	for _, kata := range kataperKata {
		switch {
		case strings.ContainsRune(kata, 'a'):
			total++
		case strings.ContainsRune(kata, 'i'):
			total++
		case strings.ContainsRune(kata, 'e'):
			total++
		case strings.ContainsRune(kata, 'o'):
			total++
		case strings.ContainsRune(kata, 'u'):
		case :
			total = total
			total++
		default:
			total = total
		}
	}
	return int32(total)
}

func main() {
	s := "aku aku aja sosi$"
	fmt.Println(countValidWords(s))
}
