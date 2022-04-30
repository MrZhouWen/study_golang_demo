package 质因子_位运算

/*
给定一个正整数 n，找到并返回 n 的二进制表示中两个 相邻 1 之间的 最长距离 。如果不存在两个相邻的 1，返回 0 。
	输入：n = 5
	输出：2
	解释：5 的二进制是 "101" 。

思路：位运算
T = O(logn) 因为每次n右移，减少一半
S = O(1)
*/
func binaryGap(n int) (ans int) {
	for i, last := 0, -1; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				ans = max(ans, i-last)
			}
			last = i
		}
		n >>= 1
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
