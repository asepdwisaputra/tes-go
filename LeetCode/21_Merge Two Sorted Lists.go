/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	wadah := []int{}

	for list1 != nil {
		wadah = append(wadah, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		wadah = append(wadah, list2.Val)
		list2 = list2.Next
	}

	fmt.Println(wadah)

	sort.Ints(wadah)

	fmt.Println(wadah)

	var head, tail *ListNode
	for _, val := range wadah {
		node := &ListNode{Val: val, Next: nil}
		if head == nil {
			head = node
			tail = node
			//fmt.Println(head)
		} else {
			tail.Next = node
			tail = node
			//fmt.Println(tail)
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
	fmt.Println(mergeTwoLists(l1, l2))
}
