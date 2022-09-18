package reverselist

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	curr := head
	next := curr.Next
	var prev *ListNode
	for {
		curr.Next = prev
		prev = curr
		curr = next
		if curr == nil {
			break
		}
		next = curr.Next
	}

	return prev
}

func Test() {
	head := &ListNode{Val: 1}
	curr := head
	for i := 2; i <= 5; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}
	for curr = reverseList(head); curr != nil; curr = curr.Next {
		fmt.Println(curr.Val)
	}
}
