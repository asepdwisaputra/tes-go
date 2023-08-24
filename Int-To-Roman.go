package main

import (
	"fmt"
	"strings"
)

var encode = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	var result strings.Builder

	//fmt.Println("!", encode)
	for _, symbol := range encode {
		//fmt.Println(symbol)
		for num >= symbol.Value {
			result.WriteString(symbol.Symbol)
			num -= symbol.Value
		}
	}

	return result.String()
}

func main() {
	num := 14
	roman := intToRoman(num)
	fmt.Println(roman)
}
