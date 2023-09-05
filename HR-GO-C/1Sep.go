package main

import (
	// "bufio"
	"fmt"
	// "io"
	// "os"
	// "strconv"
	// "strings"
)

/*
 * Complete the 'restock' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY itemCount
 *  2. INTEGER target
 */

func restock(itemCount []int32, target int32) int32 {
	// Write your code here
	var temp int32
	var keluar int32
	for _, nilai := range itemCount {
		if temp == target {
			keluar = 0
			break
		}
		if temp > int32(target) {
			keluar = temp - target
		}
		if temp < target {
			temp += int32(nilai)
		}
	}
	// for i := 0; i < len(itemCount); i++ {
	// 	fmt.Println(i)
	// 	temp += int32(i)
	// 	if temp == target {
	// 		keluar = 0
	// 		break
	// 	}
	// 	if temp > int32(target) {
	// 		keluar = temp - target
	// 	}
	// }
	fmt.Println(temp)
	return keluar // harusnya keluar - temp
}

func main() {
	item := []int32{1, 2, 3, 4, 10}
	a := restock(item, 4)
	fmt.Println(a)
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//     checkError(err)

//     defer stdout.Close()

//     writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

//     itemCountCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var itemCount []int32

//     for i := 0; i < int(itemCountCount); i++ {
//         itemCountItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         itemCountItem := int32(itemCountItemTemp)
//         itemCount = append(itemCount, itemCountItem)
//     }

//     targetTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     target := int32(targetTemp)

//     result := restock(itemCount, target)

//     fmt.Fprintf(writer, "%d\n", result)

//     writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }
