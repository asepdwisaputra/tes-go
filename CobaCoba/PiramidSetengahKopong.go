package main

import "fmt"

func main() {
	totalBaris := 10
	bintang := "*"
	spasi := " "
	// Vertikal sebanyak totalBaris
	for i := 1; i <= totalBaris; i++ {
		// Horizontal dengan ketentuan
		for j := 1; j <= i; j++ {
			if i == totalBaris { // Jika baris terakhir -> full bintang
				fmt.Print(bintang)
			} else if j == 1 || j == i { // Jika panggilan ke-1 atau terakhir -> bintang
				fmt.Print(bintang)
			} else {
				fmt.Print(spasi) // Panggilan lainnya -> Kosong
			}
		}
		fmt.Println("") // Ganti baris
	}
}

/*RESULT

*
**
* *
*  *
*   *
*    *
*     *
*      *
*       *
**********

 */
