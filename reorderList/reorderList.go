package reorderlist

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：

// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) *ListNode {
	count := 0
	var curr *ListNode

	for curr = head; curr != nil; curr = curr.Next {
		count++
	}
	stack := make([]*ListNode, count)
	pos := 0
	for curr := head; curr != nil; curr = curr.Next {
		stack[pos] = curr
		pos++
	}

	left, right := 0, count-1
	dummy := &ListNode{}
	curr = dummy
	flag := false
	for i := 0; i < count; i++ {
		if flag {
			curr.Next = stack[right]
			right--
		} else {
			curr.Next = stack[left]
			left++
		}
		curr = curr.Next
		flag = !flag
	}
	curr.Next = nil

	return dummy.Next
}

func Test() {
	reorderList(nil)
}
