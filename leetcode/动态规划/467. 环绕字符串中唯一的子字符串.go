package 动态规划

/**
把字符串 s 看作是 “...xyzabcdefghijklmnopqrstuvwxyzabc...” 的无限环绕字符串
给定一个字符串p, 求p的非空子串中（去重的）有多少在s中的？

输入: p = "cac"
输出: 2
解释: 字符串 p 中的去重非空子串 只有两个子串“a”、“c” 在s中

思路：dp T= O(n) S = O(26)
假设 p 中的某一个字符x, 记 dp[x]表示，p中的，以x为结尾且在s中的子串的个数

参考：https://leetcode.cn/problems/unique-substrings-in-wraparound-string/solution/huan-rao-zi-fu-chuan-zhong-wei-yi-de-zi-ndvea/
*/
func findSubstringInWraproundString(p string) (ans int) {
	dp := [26]int{}
	k := 0
	for i, ch := range p {
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 { //要么是bc这种相差1，要么是za这种，相差-25
			k++
		} else {
			k = 1
		}
		dp[ch-'a'] = max(dp[ch-'a'], k) //可能先单独遍历过a,k=1;  然后又遍历到za,k变成2了
	}
	for _, v := range dp {
		ans += v
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
