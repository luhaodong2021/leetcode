package wordbreak

import (
	"fmt"
	"strings"
)

func recursion(s string, wordDict []string, l int, record map[int]bool) bool {
	if _, ok := record[l]; ok {
		return false
	}
	for _, w := range wordDict {
		if strings.HasPrefix(s, w) {
			if len(s) == len(w) {
				return true
			}
			l = l + len(w)
			if !recursion(s[len(w):], wordDict, l, record) {
				record[l] = false
				continue
			} else {
				return true
			}
		}
	}
	return false
}

func dp(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dict := make(map[string]struct{})
	for _, w := range wordDict {
		dict[w] = struct{}{}
	}

	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}
			if _, ok := dict[s[j:i]]; ok {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func wordBreak(s string, wordDict []string) bool {
	record := make(map[int]bool)
	return recursion(s, wordDict, 0, record)
}

func Test() {
	fmt.Println(wordBreak("leetcode",
		[]string{"leet", "code"}),
	)
	fmt.Println(dp("leetcode",
		[]string{"leet", "code"}),
	)
}
