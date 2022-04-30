package 模拟

/**
给你一个字符串 s 和一个字符 c ，且 c 是 s 中出现过的字符。
返回一个整数数组 answer ，其中 answer.length == s.length 且 answer[i] 是 s 中从下标 i 到离它 最近 的字符 c 的 距离 。
两个下标i 和 j 之间的 距离 为 abs(i - j) ，其中 abs 是绝对值函数。

思路：2次遍历 T= O(n) S= O(1)
问题可以转换成，对 s 的每个下标 i，求：
	1、s[i] 到其左侧最近的字符 c 的距离
	2、s[i] 到其右侧最近的字符 c 的距离
这两者的最小值。
*/
func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	//从左往右遍历时，起始的一些位置，可能还未遇到c，所以idx置为很大
	idx := -2 * n
	for i, ch := range s {
		if byte(ch) == c {
			idx = i
		}
		ans[i] = i - idx
	}
	//从右往左遍历，同理的最右边的一些起始位置，可能未遇到c，所以，将idx置为很大
	idx = 2 * n
	for i := n - 1; i >= 0; i-- {
		if byte(s[i]) == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
