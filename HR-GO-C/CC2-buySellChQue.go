package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getNetProfit(events []string) []int64 {
	hasil := []int64{}
	port := make(map[string]int64)
	var profit int64

	for _, event := range events {
		bagian := strings.Split(event, " ")
		bagianDep := bagian[0]

		switch bagianDep {
		case "BUY", "SELL", "CHANGE":
			brand := bagian[1]
			nilai := bagian[len(bagian)-1]

			switch bagianDep {
			case "BUY":
				angka, _ := strconv.ParseInt(nilai, 10, 64)
				port[brand] = port[brand] + angka

			case "SELL":
				angka, _ := strconv.ParseInt(nilai, 10, 64)
				port[brand] = port[brand] - angka

			case "CHANGE":
				angka, _ := strconv.ParseInt(nilai, 10, 64)
				if outnya, ada := port[brand]; ada {
					profit = angka * outnya
				}
			}

		case "QUERY":
			//hasil = append(hasil, profit)
			if profit >= 0 {
				hasil = append(hasil, profit)
			} else {
				hasil = append(hasil, profit)
				hasil[len(hasil)-1] = hasil[len(hasil)-1] + hasil[len(hasil)-2]
			}
		}
	}
	return hasil
}

func main() {
	//slice := []string{"BUY googl 20", "BUY aapl 50", "CHANGE googl 6", "QUERY"} //120
	slice := []string{"BUY googl 20", "BUY aapl 50", "CHANGE googl 6", "QUERY", "SELL aapl 10", "CHANGE aapl -2", "QUERY" /*, "CHANGE googl -2", "QUERY"*/} //120-40

	fmt.Println(getNetProfit(slice))
}
