package reversebetween

import "fmt"

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
// 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var (
		nodeBeforeReverse *ListNode
		nodeAfterReverse  *ListNode
		nodeStartReverse  *ListNode
		nodeEndReverse    *ListNode
	)
	var curr, prev, next *ListNode

	sn := 1
	curr = head
	next = head.Next
	for {
		switch {
		case sn < left:
		case sn == left:
			nodeBeforeReverse = prev
			nodeStartReverse = curr
			curr.Next = prev
			fallthrough
		case sn == right:
			nodeAfterReverse = next
			nodeEndReverse = curr
			curr.Next = prev
		case sn > right:
			if nodeBeforeReverse != nil {
				nodeBeforeReverse.Next = nodeEndReverse
			}
			nodeStartReverse.Next = nodeAfterReverse
			if left == 1 {
				return nodeEndReverse
			}
			return head
		default:
			curr.Next = prev
		}
		prev = curr
		curr = next
		if curr != nil {
			next = curr.Next
		}
		sn++
	}
}

func Test() {
	head := &ListNode{Val: 1}
	curr := head
	for i := 2; i <= 5; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}
	head = reverseBetween(head, 1, 4)
	for curr = head; curr != nil; curr = curr.Next {
		fmt.Println(curr.Val)
	}
}
