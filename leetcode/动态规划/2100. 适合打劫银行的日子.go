package 动态规划

/**
题目：https://leetcode.cn/problems/find-good-days-to-rob-the-bank/

思路：DP
动态规划 预处理每个数字的 左测依次递减的数量 lefti 右侧依次递增的数量 righti
T(n) S(n)
*/
func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	left := make([]int, n)
	right := make([]int, n)
	//预处理每一个数左右2测的递减和递增数量
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
		if security[n-i-1] <= security[n-i] {
			right[n-i-1] = right[n-i] + 1
		}
	}

	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}

	return
}
