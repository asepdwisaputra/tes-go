package main

import (
	"fmt"
	"math"
)

func akar(n int) int {
	for i := 1; i > 0; i++ {
		if i*i == n {
			return int(i)
		}
	}
	return 0
}

func akar2(n float64) float64 {
	return math.Sqrt(n)
}

func main() {
	fmt.Println(akar(25))
	fmt.Println(akar(9))
	fmt.Println(akar2(25))
	fmt.Println(akar2(9))

}
