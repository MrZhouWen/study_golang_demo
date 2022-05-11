package 贪心

/**
由范围 [0,n] 内所有整数组成的 n + 1 个整数的排列序列可以表示为长度为 n 的字符串 s ，其中:
	如果 perm[i] < perm[i + 1]，那么 s[i] == 'I'
	如果 perm[i] > perm[i + 1]，那么 s[i] == 'D'
给定一个字符串 s ，重构排列perm 并返回它。如果有多个有效排列perm，则返回其中 任何一个 。

输入：s = "IDID"
输出：[0,4,1,3,2]

思路：贪心 T= O(n) S = O(1)
考虑 perm[0] 的值，根据题意：
	如果 s[0]=‘I’，那么令 perm[0]=0，则无论 perm[1] 为何值都满足 perm[0]<perm[1]；
	如果 s[0]=‘D’，那么令 perm[0]=n，则无论 perm[1] 为何值都满足 perm[0]>perm[1]；
确定好 perm[0] 后，剩余的 n−1 个字符和 n 个待确定的数就变成了一个和原问题相同，但规模为 n−1 的问题。

因此我们可以继续按照上述方法确定 perm[1]：
	如果 s[1]=‘I’，那么令 perm[1] 为剩余数字中的最小数；
	如果 s[1]=‘D’，那么令 perm[1] 为剩余数字中的最大数。
如此循环直至剩下一个数，填入 perm[n] 中。
*/
func diStringMatch(s string) []int {
	n := len(s)
	perm := make([]int, n+1)
	low, high := 0, n
	for i, ch := range s {
		if ch == 'I' {
			perm[i] = low
			low++
		} else if ch == 'D' {
			perm[i] = high
			high--
		}
	}
	perm[n] = low //只剩最后一个的时候，low == high
	return perm
}
