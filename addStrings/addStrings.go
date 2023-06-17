package addstrings

import "fmt"

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

func addStrings(num1 string, num2 string) string {
	var rv []byte
	if len(num1) > len(num2) {
		rv = make([]byte, len(num1)+1)
	} else {
		rv = make([]byte, len(num2)+1)
	}
	reverseIndex := len(rv) - 1
	carry := byte(0)

	var bit byte
	for len(num1) > 0 || len(num2) > 0 {
		bit = 0
		if len(num1) > 0 {
			bit = num1[len(num1)-1] - '0'
			num1 = num1[:len(num1)-1]
		}
		if len(num2) > 0 {
			bit = bit + num2[len(num2)-1] - '0'
			num2 = num2[:len(num2)-1]
		}
		bit := bit + carry
		if bit < 10 {
			carry = 0
		} else {
			bit = bit - 10
			carry = 1
		}
		rv[reverseIndex] = bit + '0'
		reverseIndex--
	}
	if carry == 1 {
		rv[0] = '1'
		return string(rv)
	}
	return string(rv[1:])
}

func Test() {
	fmt.Println(addStrings("1", "99"))
}
