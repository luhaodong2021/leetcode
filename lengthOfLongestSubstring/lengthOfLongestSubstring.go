package lengthoflongestsubstring

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// hashmap
func lengthOfLongestSubstring(s string) int {
	duplicate := make(map[byte]int)

	var retval int
	var i int
	var startOfSubString int

	for i = 0; i < len(s); i++ {
		c := s[i]
		position, exist := duplicate[c]
		if !exist {
			duplicate[c] = i
			continue
		}
		if i-startOfSubString > retval {
			retval = i - startOfSubString
		}
		for j := startOfSubString; j <= position; j++ {
			delete(duplicate, s[j])
		}
		duplicate[c] = i
		startOfSubString = position + 1
	}
	if i-startOfSubString > retval {
		retval = i - startOfSubString
	}
	return retval
}

// array
func lengthOfLongestSubstring1(s string) int {
	duplicate := [128]uint16{65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535, 65535}

	var retval uint16
	var i uint16
	var startOfSubString uint16

	for i = 0; i < uint16(len(s)); i++ {
		c := s[i]
		position := duplicate[c]
		if position == 65535 {
			duplicate[c] = i
			continue
		}
		if i-startOfSubString > retval {
			retval = i - startOfSubString
		}
		for j := startOfSubString; j <= position; j++ {
			duplicate[(s[j])] = 65535
		}
		duplicate[c] = i
		startOfSubString = position + 1
	}
	if i-startOfSubString > retval {
		retval = i - startOfSubString
	}
	return int(retval)
}
