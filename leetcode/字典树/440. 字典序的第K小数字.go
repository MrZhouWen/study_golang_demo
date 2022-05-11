package 字典树

/**
题目：
给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。

思路：
字典树 其实不需要构造字典树，因为计算有规律 T((logn)^2) S(1)
*/
func findKthNumber(n, k int) int {
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if steps < k+1 { //说明结果不在这课子树中
			k -= steps
			cur++
		} else { //就在这个字典树里找
			k--
			cur *= 10
		}
	}
	return cur
}

func getSteps(cur, n int) (steps int) { //层序遍历就可以获得，以cur为根的子树的所有节点数
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first = first * 10
		last = last*10 + 9
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
