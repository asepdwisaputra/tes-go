package main

import (
	"fmt"
	"sort"
)

func greedyCoinChange(coins []int, target int) []int {
	// Urutkan koin dalam urutan menurun(Besar-Kecil).
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	fmt.Println(coins)
	coinCount := make([]int, len(coins))

	for i, coin := range coins {
		for target >= coin {
			target -= coin
			coinCount[i]++
		}
	}

	return coinCount
}

func main() {
	availableCoins := []int{25, 10, 5, 1, 100} // Set koin yang tersedia (urutan menurun).
	amount := 42

	coinCounts := greedyCoinChange(availableCoins, amount)

	fmt.Printf("Koin yang digunakan untuk menghasilkan %d: \n", amount)
	for i, count := range coinCounts {
		if count > 0 {
			fmt.Printf("%d koin(s) x %d\n", availableCoins[i], count)
		}
	}

	/*
		Coba bayangkan apabila availableCoins = 3, 5, 7
		lalu amount = 11
		
		Greedy akan error
	/
}
