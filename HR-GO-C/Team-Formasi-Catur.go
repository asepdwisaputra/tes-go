package main

import "fmt"

func countTeams(rating []int32, queries [][]int32) /*[]int32*/ {
	//grup := make(map[int32]int32)
	// for _, v := range rating {
	// 	grup[v]++
	// }
	// for i := 0; i < len(rating); i++ {
	// 	for j := i + 1; j < len(rating); j++ {
	// 		if rating[i] == rating[j]{
	// 			queries = append(queries, int32(rating[i], rating[j])
	// 			continue
	// 		}
	// 	}
	// }
	//fmt.Println(grup)
}

func main() {
	rating := []int32{1, 2, 2, 3}
	q := [][]int32{{1, 2}, {1, 1}}

	countTeams(rating, q)

	// a := countTeams(rating, q)
	// fmt.Println(a)
}
