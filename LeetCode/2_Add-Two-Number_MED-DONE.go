package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addLinkedLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		fmt.Println(carry)
		current.Next = &ListNode{Val: sum % 10}
		fmt.Println(current.Next)
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
		fmt.Println(current.Next)
	}
	fmt.Println(dummy.Next)
	return dummy.Next
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Membuat linked list pertama: 2 -> 4 -> 3
	list1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}

	// Membuat linked list kedua: 5 -> 6 -> 4
	list2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	// Menambahkan dua linked list
	result := addLinkedLists(list1, list2)

	// Mencetak hasilnya
	fmt.Println("Hasil Penambahan:")
	printList(result)
}
