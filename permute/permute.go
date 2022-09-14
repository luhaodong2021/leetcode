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

func permute(nums []int) [][]int {
	var path []int
	used := make([]int, len(nums))
	dfs(nums, path, used)
	return ans
}

func Test() {
	rv := permute([]int{5, 4, 6, 1})
	fmt.Println(rv)
}
