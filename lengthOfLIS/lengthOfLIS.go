package lengthoflis

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	lengthOfLIS := 1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		if dp[i] > lengthOfLIS {
			lengthOfLIS = dp[i]
		}
	}
	return lengthOfLIS
}

func Test() {
	lengthOfLIS(nil)
}
