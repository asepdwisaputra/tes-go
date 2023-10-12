package main

import "fmt"

func main() {
	/*
		Buatlah sebuah program yang mencetak angka dari 1 sampai 100
		dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka,
		dan untuk kelipatan '5' cetak "Buzz”.
		Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz …….
	*/
	angka := 100
	for i := 1; i <= angka; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print( /*i,*/ "Fizz Buzz")
			fmt.Print(", ")
		case i%3 == 0:
			fmt.Print( /*i,*/ "Fizz") // agar mudah dibaca saya masukin i depan Fizz
			fmt.Print(", ")
		case i%5 == 0:
			fmt.Print( /*i,*/ "Buzz") // Jika case diatas true ini tak jalan
			fmt.Print(", ")

		default:
			fmt.Print(i)
			fmt.Print(", ")
		}
	}
}
