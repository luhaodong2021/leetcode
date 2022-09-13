package coinchange

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dict := make(map[int]struct{})
	for _, c := range coins {
		dict[c] = struct{}{}
	}
	for i := 1; i <= amount; i++ {
		tmp := i + 1
		for j := 0; j < i; j++ {
			if dp[j] == -1 {
				continue
			}
			if _, ok := dict[i-j]; ok && dp[j]+1 < tmp {
				tmp = dp[j] + 1
			}
		}
		if tmp != i+1 {
			dp[i] = tmp
		} else {
			dp[i] = -1
		}
	}
	return dp[amount]
}

func Test() {
	fmt.Println(coinChange([]int{80, 149, 71, 111}, 8683))
}
