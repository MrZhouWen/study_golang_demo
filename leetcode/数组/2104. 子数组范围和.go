package 数组

/**
给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。
返回 nums 中 所有 子数组范围的 和 。
子数组是数组中一个连续 非空 的元素序列。

思路：单调栈 T = O(n) S = O(n)
求出以每一个nums[i]为最大值 或 最小值，的所有子数组的，sumMax - sumMin
参考：https://leetcode.cn/problems/sum-of-subarray-ranges/solution/zi-shu-zu-fan-wei-he-by-leetcode-solutio-lamr/
*/
/**
单调栈 T(n) S(n)
*/
func subArrayRanges(nums []int) int64 {
	n := len(nums)
	minLeft := make([]int, n) //记录每一个元素nums[i]的左边最近的 <= nums[i]的下标
	maxLeft := make([]int, n) //记录每一个元素nums[i]的左边最近的 > nums[i]的下标
	var minStk, maxStk []int
	for i, num := range nums {
		for len(minStk) > 0 && nums[minStk[len(minStk)-1]] > num {
			minStk = minStk[:len(minStk)-1]
		}
		if len(minStk) > 0 {
			minLeft[i] = minStk[len(minStk)-1]
		} else {
			minLeft[i] = -1
		}
		minStk = append(minStk, i)

		// 如果 nums[maxStk[len(maxStk)-1]] == num, 那么根据定义，
		// nums[maxStk[len(maxStk)-1]] 逻辑上小于 num，因为 maxStk[len(maxStk)-1] < i
		for len(maxStk) > 0 && nums[maxStk[len(maxStk)-1]] <= num {
			maxStk = maxStk[:len(maxStk)-1]
		}
		if len(maxStk) > 0 {
			maxLeft[i] = maxStk[len(maxStk)-1]
		} else {
			maxLeft[i] = -1
		}
		maxStk = append(maxStk, i)
	}

	minRight := make([]int, n)
	maxRight := make([]int, n)
	minStk = minStk[:0]
	maxStk = maxStk[:0]
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		// 如果 nums[minStk[len(minStk)-1]] == num, 那么根据定义，
		// nums[minStk[len(minStk)-1]] 逻辑上大于 num，因为 minStk[len(minStk)-1] > i
		for len(minStk) > 0 && nums[minStk[len(minStk)-1]] >= num {
			minStk = minStk[:len(minStk)-1]
		}
		if len(minStk) > 0 {
			minRight[i] = minStk[len(minStk)-1]
		} else {
			minRight[i] = n
		}
		minStk = append(minStk, i)

		for len(maxStk) > 0 && nums[maxStk[len(maxStk)-1]] < num {
			maxStk = maxStk[:len(maxStk)-1]
		}
		if len(maxStk) > 0 {
			maxRight[i] = maxStk[len(maxStk)-1]
		} else {
			maxRight[i] = n
		}
		maxStk = append(maxStk, i)
	}

	var sumMax, sumMin int64
	for i, num := range nums {
		sumMax += int64(maxRight[i]-i) * int64(i-maxLeft[i]) * int64(num)
		sumMin += int64(minRight[i]-i) * int64(i-minLeft[i]) * int64(num)
	}
	return sumMax - sumMin
}
