package main

import "fmt"

func Pangkat(dasar, pangkat int) int {
	nilai := 1
	for i := 1; i <= pangkat; i++ {
		nilai = nilai * dasar
	}
	return nilai
}

func main() {
	nilaiPangkat := 3
	nilaiDasar := 3

	fmt.Println(Pangkat(nilaiDasar, nilaiPangkat))
}
