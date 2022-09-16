package longestpalindrome

import "fmt"

// 最长回文子串

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			if isPalindrome(s[j : j+i]) {
				return s[j : j+i]
			}
		}
	}
	return ""
}

func Test() {
	fmt.Println(longestPalindrome("cccc"))
}
