package maxsubarray

import "fmt"

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组 是数组中的一个连续部分。

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	rv := dp[0]
	if dp[0] < 0 {
		dp[0] = 0
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] + nums[i]
		if dp[i] > rv {
			rv = dp[i]
		}
		if dp[i] < 0 {
			dp[i] = 0
		}
	}
	return rv
}

func Test() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //  sum[4,-1,2,1] = 6
	fmt.Println(maxSubArray([]int{-2, -1, -3}))
}
