package 数组

import "sort"

/**
统计一个数字在排好序数组中出现的次数。

思路：二分查找
T = O(logn) S = O(1)
*/
func search(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return rightmost - leftmost + 1
}
