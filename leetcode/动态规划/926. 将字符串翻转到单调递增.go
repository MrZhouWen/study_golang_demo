package 动态规划

/**
将一个二进制字符串 转换为 递增的字符串 需要翻转的 最少次数

输入：s = "00110"
输出：1
解释：翻转最后一位得到 00111

思路： DP T= O(n) S= O(1)
用 dp[i][0] 和 dp[i][1] 分别表示 翻转调整后(可能不需要翻转) 下标 i 处的字符为 0 和 1 的情况下使得 s[0..i] 单调递增的最小翻转次数
	dp[i][0] = dp[i−1][0] + I(s[i]=‘1’)
	dp[i][1] = min(dp[i−1][0],dp[i−1][1]) + I(s[i]=‘0’)， 其中I(表达式)为示性函数，表达式为true返回1， 为假返回0
初始，可以定义 dp[-1][0] = dp[-1][1] = 0, 那么 i  可以，从下标0开始算
*/
func minFlipsMonoIncr(s string) int {
	dp0, dp1 := 0, 0 //此处利用滚动数组将S= O(1) => dp0 = dp[i-1][0] ; dp1 = dp[i-1][1]
	for _, c := range s {
		dp0New, dp1New := dp0, min(dp0, dp1)
		if c == '1' {
			dp0New++
		} else {
			dp1New++
		}
		dp0, dp1 = dp0New, dp1New
	}
	return min(dp0, dp1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
