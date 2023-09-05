/*
	Pendekatan top-down mulai dari masalah besar dan
	memecahkannya menjadi submasalah
*/

package main

import (
	"fmt"
)

var memo map[int]int

func fibonacci(n int) int {
	// Memeriksa apakah hasil sudah ada dalam memo
	if val, ok := memo[n]; ok {
		return val
	}

	// Basis: Fibonacci dari 0 dan 1 adalah diri mereka sendiri
	if n <= 1 {
		return n
	}

	// Rekursif: Menghitung Fibonacci menggunakan rekursi
	result := fibonacci(n-1) + fibonacci(n-2)

	// Menyimpan hasil perhitungan dalam memo
	memo[n] = result
	//fmt.Println(memo)

	return result
}

func main() {
	n := 10
	memo = make(map[int]int)

	result := fibonacci(n)
	fmt.Printf("Fibonacci ke-%d adalah %d\n", n, result)
	//2,186s
}
