package main

import "fmt"

func main() {
	flagHeight := 5
	bigJump := 2

	tbigJump := flagHeight / bigJump

	var total int
	if flagHeight > 0 {
		total = 0
	}

	if flagHeight < bigJump {
		total = 1
	}

	if flagHeight%bigJump == 0 {
		total = total + tbigJump
	} else {
		sisa := flagHeight % bigJump
		total = sisa + tbigJump
	}

	fmt.Println(total)
}
