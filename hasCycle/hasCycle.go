package hascycle

// 给你一个链表的头节点 head ，判断链表中是否有环。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针,有环快指针必然追上慢指针
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	slowShouldMove := false
	for fast != nil {
		fast = fast.Next
		if slowShouldMove {
			slow = slow.Next
		}
		if fast == slow {
			return true
		}
		slowShouldMove = !slowShouldMove
	}
	return false
}

func Test() {
	hasCycle(nil)
}
