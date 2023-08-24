package main

import (
	"fmt"
)

var decode = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	first := decode[s[0]]
	if len(s) == 1 {
		return first
	}

	next := decode[s[1]]
	if next > first { // IV
		return (next - first) + romanToInt(s[2:])
	}
	return first + romanToInt(s[1:])
}

func main() {
	s := "XIII"

	fmt.Println(decode[s[2]])

	// fmt.Println(romanToInt(s))
}
