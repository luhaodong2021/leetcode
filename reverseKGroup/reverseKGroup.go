package reversekgroup

import "fmt"

// Definition for singly-linked list.
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

func listTail(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for {
		if head.Next == nil {
			return head
		}
		head = head.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	sn := 0
	l := head
	groupHead := l
	n := l
	groupTail := &ListNode{}
	rv := groupTail
	for l != nil {
		sn++
		if sn == k {
			n = l.Next
			l.Next = nil
			newList := reverseList(groupHead)
			groupTail.Next = newList
			groupTail = listTail(newList)
			sn = 0
			l = n
			groupHead = l
			continue
		}
		l = l.Next
	}
	if sn < k {
		groupTail.Next = n
	}
	return rv.Next
}

func Test() {
	head := &ListNode{Val: 1}
	curr := head
	for i := 2; i <= 10; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}
	for curr = reverseKGroup(head, 3); curr != nil; curr = curr.Next {
		fmt.Println(curr.Val)
	}
}
