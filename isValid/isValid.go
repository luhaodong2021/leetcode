package isvalid

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

func isValid(s string) bool {
	stack := make([]byte, len(s)/2)
	pos := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ')':
			pos--
			if pos < 0 {
				return false
			}
			if stack[pos] == '(' {
				continue
			}
			return false
		case ']':
			pos--
			if pos < 0 {
				return false
			}
			if stack[pos] == '[' {
				continue
			}
			return false
		case '}':
			pos--
			if pos < 0 {
				return false
			}
			if stack[pos] == '{' {
				continue
			}
			return false

		default:
			if pos == len(stack) {
				return false
			}
			stack[pos] = s[i]
			pos++
		}
	}
	return pos == 0
}

func Test() {
	fmt.Println(isValid("{{}[]()}"))
}
