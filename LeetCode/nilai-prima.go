package main

import "fmt"

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	//Cara 1
	// for i := 2; i*i <= number; i++ { // 2,3,4,5,6,7
	// 	if number%i == 0 {
	// 		return false
	// 	}
	// }

	//Cara 2
	if number == 2 {
		return true
	}
	if number%2 == 0 {
		return false
	}
	// 5,7 = 35
	for i := 3; i*i <= number; i += 2 { //harus pake =
		// karena jika i*i < number -- untuk nilai primeNumber(9)
		// i+=2 menyingkat angka genap
		if number%i == 0 {
			return false
		}
	}

	////////////////////////
	return true
}

func main() {
	fmt.Println(primeNumber(7))
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
