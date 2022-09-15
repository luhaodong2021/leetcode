package permute

import "fmt"

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
var ans [][]int

func dfs(nums []int, path []int, used []int) {
	if len(path) == len(nums) {
		item := make([]int, len(nums))
		copy(item, path)
		ans = append(ans, item)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] == 1 {
			continue
		}
		used[i] = 1
		path = append(path, nums[i])
		dfs(nums, path, used)
		used[i] = 0
		path = path[:len(path)-1]
	}
}

type Node struct {
	path []int
	used []int
}

func bfs(node Node, nums []int) {
	if len(node.path) == len(nums) {
		record := make([]int, len(nums))
		copy(record, node.path)
		ans = append(ans, record)
	}
	for i := 0; i < len(nums); i++ {
		if node.used[i] == 1 {
			continue
		}
		newNode := Node{make([]int, len(node.path)), make([]int, len(nums))}
		copy(newNode.path, node.path)
		copy(newNode.used, node.used)
		newNode.path = append(newNode.path, nums[i])
		newNode.used[i] = 1
		bfs(newNode, nums)
	}
}

func permute(nums []int) [][]int {
	// dfs
	var path []int
	used := make([]int, len(nums))
	dfs(nums, path, used)
	// bfs
	newNode := Node{used: make([]int, len(nums))}
	bfs(newNode, nums)
	return ans
}

func Test() {
	rv := permute([]int{5, 4, 6, 1})
	fmt.Println(rv)
}
