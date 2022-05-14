package 数组

import "sort"

/**
https://leetcode.cn/problems/B1IidL/
*/

/*
思路：二分查找 T(logn) S(1)
*/
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool {
		return arr[i] > arr[i+1]
	})
}
