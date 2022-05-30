package 数组

import "sort"

/**
给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。
在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。
*/
/*
排序：将数组排序后，首尾元素依次配对，组成一个个区间，当x落在所有区间的时候，所需的步数最少
    当 n = len(nums) 为偶数，中间的区间是 nums[n/2-1], nums[n/2]
    当 n = len(nums) 为奇数，中间的 数是  nums[n/2]，
    所以，我们取折中方案，x = nums[n/2]满足题意
T = O(nlogn) S= O(logn)
*/
func minMoves2(nums []int) (ans int) {
	sort.Ints(nums)
	x := nums[len(nums)/2]
	for _, num := range nums {
		ans += abs(num - x)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
