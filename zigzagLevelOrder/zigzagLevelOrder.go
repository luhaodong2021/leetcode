package zigzaglevelorder

import "container/list"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	left2Right := true
	res := make([][]int, 0)
	cache := list.New()
	if root != nil {
		cache.PushFront(root)
	}

	for cache.Len() > 0 {
		level := make([]int, cache.Len())
		for size := cache.Len(); size > 0; size-- {
			root = cache.Front().Value.(*TreeNode)
			if left2Right {
				level[len(level)-size] = root.Val
			} else {
				level[size-1] = root.Val
			}
			cache.Remove(cache.Front())

			if root.Left != nil {
				cache.PushBack(root.Left)
			}
			if root.Right != nil {
				cache.PushBack(root.Right)
			}
		}
		res = append(res, level)
		left2Right = !left2Right
	}
	return res
}

func Test() {
	zigzagLevelOrder(nil)
}
