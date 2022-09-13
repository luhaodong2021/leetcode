package reverselist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	curr := head
	next := curr.Next
	var prev *ListNode
	for next != nil {
		prev = curr
		curr = next
		next = curr.Next
		curr.Next = prev
	}

	return curr
}

func Test() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	r := reverseList(head)
	fmt.Println(r.Val, r.Next.Val, r.Next.Next.Val)
}
