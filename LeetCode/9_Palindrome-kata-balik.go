package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	// Balikan Input
	xStr := strconv.Itoa(x) // Jadi string
	balikanStr := ""
	for i := len(xStr) - 1; i >= 0; i-- {
		balikanStr += string(xStr[i])
		fmt.Println(balikanStr)
	}
	balikanInt, _ := strconv.Atoi(balikanStr)
	if x == balikanInt {
		return true
	} else {
		return false
	}
}

func main() {
	x := 12321
	fmt.Println(isPalindrome(x))
}
