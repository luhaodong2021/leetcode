package mergetwolists

// 合并两个有序链表

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	rv := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			head = head.Next
			list1 = list1.Next
		} else {
			head.Next = list2
			head = head.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return rv.Next
}

func Test() {
	mergeTwoLists(nil, nil)
}
