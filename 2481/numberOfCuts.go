package numberOfCuts

import "fmt"

// 圆内一个 有效切割 符合以下二者之一

// 该切割是两个端点在圆上的线段 且该线段经过圆心。
// 该切割是一端在圆心另一端在圆上的线段。

// 给你一个整数 n 请你返回将圆切割成相等的 n 等分的 最少 切割次数

func numberOfCuts(n int) int {
	if n ==1 {
		return 0;
	}

	odd := n&0x01 == 1

	if odd {
		return n
	}

	return n >> 1
}

func Test() {
	fmt.Println(numberOfCuts(2))
}
