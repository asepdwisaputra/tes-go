package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	slice1 := []int{}
	slice2 := []int{}

	// masukan data ke Slice
	for l1 != nil {
		slice1 = append(slice1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		slice2 = append(slice2, l2.Val)
		l2 = l2.Next
	}

	// Balikan Slice
	wadah1 := strings.Builder{}
	for i := len(slice1) - 1; i >= 0; i-- {
		wadah1.WriteString(strconv.Itoa(slice1[i]))
	}
	wadah2 := strings.Builder{}
	for j := len(slice2) - 1; j >= 0; j-- {
		wadah2.WriteString(strconv.Itoa(slice2[j]))
	}

	//Keluarkan dari wadah
	cetakan1 := wadah1.String()
	cetakan2 := wadah2.String()
	fmt.Println(cetakan1)
	fmt.Println(cetakan2)

	//Ganti tipe data dan jumlahkan
	ubah1, _ := strconv.ParseUint(cetakan1, 10, 64)
	ubah2, _ := strconv.ParseUint(cetakan2, 10, 64)
	hasil := ubah1 + ubah2
	fmt.Println(hasil)

	// hasil pecah dan jadikan slice
	hasilPecah := strconv.Itoa(int(hasil))
	hasilSlice := strings.Split(hasilPecah, "")
	fmt.Println(hasilSlice)

	//balikkan nilai hasilSlice
	hasilSliceBalik := []int{}
	for s := len(hasilSlice) - 1; s >= 0; s-- {
		sfix, _ := strconv.Atoi(hasilSlice[s])
		hasilSliceBalik = append(hasilSliceBalik, sfix)
	}
	fmt.Println(hasilSliceBalik)

	// Inisialisasi head dan tail
	var head, tail *ListNode

	// Iterasi melalui slice dan buat node untuk setiap elemen
	for _, val := range hasilSliceBalik {
		//nilai, _ := strconv.Atoi(val)
		node := &ListNode{Val: val, Next: nil}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return head
}

func main() {
	// Creating nodes for the linked list
	node1 := &ListNode{Val: 2, Next: nil}
	node2 := &ListNode{Val: 4, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}

	// Building the linked list
	node1.Next = node2
	node2.Next = node3

	// Head of the linked list
	l1 := node1

	//----------------------------------------------------------------

	// Creating nodes for the linked list
	node4 := &ListNode{Val: 5, Next: nil}
	node5 := &ListNode{Val: 6, Next: nil}
	node6 := &ListNode{Val: 4, Next: nil}

	// Building the linked list
	node4.Next = node5
	node5.Next = node6

	// Head of the linked list
	l2 := node4

	//----------------------------------------------------------------
	fmt.Println(addTwoNumbers(l1, l2))
}

// Kendala Program ini di baris 52 tampaknya data konfersi nilainya terlalu banyak
