/*
pendekatan bottom-up memecahkan submasalah terkecil terlebih dahulu
dan menggabungkannya untuk memecahkan masalah yang lebih besar.
*/
package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	// Inisialisasi array untuk menyimpan hasil Fibonacci
	fib := make([]int, n+1) // n+1 --> fib [0] s/d fib [n] .. pake var fib []int lebih enak padahal ya
	fib[0] = 0              //  kalo pake var fib []int, harus append ya, gabisa gini.
	fib[1] = 1

	// Menghitung Fibonacci dari 2 hingga n
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		//fmt.Println(fib[i])
	}

	return fib[n]
}

func main() {
	n := 10
	result := fibonacci(n)
	fmt.Printf("Fibonacci ke-%d adalah %d\n", n, result)
}

//2,151s
