package jump

func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		tmp := i
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				if dp[j]+1 < tmp {
					tmp = dp[j] + 1
				}
			}
		}
		dp[i] = tmp
	}
	// fmt.Println(dp)
	return dp[len(nums)-1]
}

func Test() {
	jump([]int{1, 2, 1, 1, 1})
}
