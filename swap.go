/*
Buat nilai variabel a dan b tertukar, bahkan dalam kasus jika nilai awalnya berubah
*DILARANG: membuat vaiabel baru ❌
*/
package main

import "fmt"

func test() {
	a := 1
	b := 2

	fmt.Println("a awal", a) // 1
	fmt.Println("b awal", b) // 2

	//StartHere
	a = a - b
	b = a + b
	a = b - a

	fmt.Println("a akhir", a) // 2
	fmt.Println("b akhir", b) // 1
}
func main() {
	test()
	// cobalah ubah variabel a pada func test menjadi 6
}
