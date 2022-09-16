package lowestcommonancestor

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(path []*TreeNode, target *TreeNode, record *[]*TreeNode) {
	curr := path[len(path)-1]
	if curr == target {
		*record = make([]*TreeNode, len(path))
		copy(*record, path)
		return
	}
	if curr.Left != nil {
		npath := append(path, curr.Left)
		dfs(npath, target, record)
	}
	if curr.Right != nil {
		npath := append(path, curr.Right)
		dfs(npath, target, record)
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var recordP, recordQ []*TreeNode
	path := []*TreeNode{root}
	dfs(path, p, &recordP)
	path = path[:1]
	dfs(path, q, &recordQ)
	lenp, lenq := len(recordP), len(recordQ)
	for i := lenp - 1; i >= 0; i-- {
		for j := 0; j < lenq; j++ {
			if recordQ[j] == recordP[i] {
				return recordP[i]
			}
		}
	}
	return nil
}

func Test() {
	lowestCommonAncestor(nil, nil, nil)
}
