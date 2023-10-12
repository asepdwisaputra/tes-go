// Algoritma Kadane => https://www.youtube.com/watch?v=GrNSGC8Z2T0
// menentukan jumlah maksimal di dalam larik

package main

import "fmt"

func jumlahMax(larik []int) int {
	nilaiMax := larik[0]
	nilaiSejauhIni := larik[0]

	for i := 1; i < len(larik); i++ {
		nilaiSejauhIni = max(larik[i], larik[i]+nilaiSejauhIni)
		nilaiMax = max(nilaiMax, nilaiSejauhIni)
	}
	return nilaiMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	larik := []int{-1, -2, 0, 1, 5, -2} // jumlah maks--> 0, 1, 5 == 6

	fmt.Println(jumlahMax(larik))
}
