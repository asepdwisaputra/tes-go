package main

import "fmt"

// Struct for a single linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to print the elements of the linked list
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Creating nodes for the linked list
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}

	// Building the linked list
	node1.Next = node2
	node2.Next = node3

	// Head of the linked list
	head := node1

	// Printing the linked list
	fmt.Println("Linked List:")
	printList(head)
}
