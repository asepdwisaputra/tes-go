/*
	Terdapat dua bilangan integer yaitu x dan n. Buatlah sebuah fungsi untuk
	melakukan perhitungan perpangkatan (x^n dibaca x pangkat n). Time complexity
	dari sebuah fungsi perpangkatan harus lebih cepat dari O(n). Contoh time
	complexity yang optimal adalah logaritmik.

	Sample Test Cases
		Input : x = 2, n = 3
		Output : 8
		Input : x = 7, n = 2
		Output : 49
*/

package main

import (
	"fmt"
	_ "math"
)

func pow(x, n int) int {
	//Cara 1
	//return int(math.Pow(float64(x), float64(n)))

	//Cara 2
	total := 1
	for i := 1; i <= n; i++ {
		total = total * x
	}
	return total
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
