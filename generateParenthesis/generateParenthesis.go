package generateparenthesis

import "fmt"

func bfs(leftCount, rightCount int, depth int, pair int, path string, res *[]string) {
	if depth == pair*2 {
		*res = append(*res, path)
		return
	}
	if leftCount == pair {
		newPath := path + ")"
		bfs(leftCount, rightCount+1, depth+1, pair, newPath, res)
		return
	}
	if leftCount == rightCount {
		newPath := path + "("
		bfs(leftCount+1, rightCount, depth+1, pair, newPath, res)
		return
	}
	newPath := path + ")"
	bfs(leftCount, rightCount+1, depth+1, pair, newPath, res)
	newPath = path + "("
	bfs(leftCount+1, rightCount, depth+1, pair, newPath, res)
}

func generateParenthesis(n int) []string {
	var ret []string
	bfs(0, 0, 0, n, "", &ret)
	return ret
}

func Test() {
	fmt.Println(generateParenthesis(2))
}
