package main

import "fmt"

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ { // 2,3,4,5,6,7
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(7))
	// fmt.Println(primeNumber(1000000007)) // true
	// fmt.Println(primeNumber(13))         // true
	// fmt.Println(primeNumber(17))         // true
	// fmt.Println(primeNumber(20))         // false
	// fmt.Println(primeNumber(35))         // false
}
