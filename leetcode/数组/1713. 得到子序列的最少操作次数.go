package 数组

import "sort"

/*
题目：
给你一个数组target，包含若干 互不相同的整数，以及另一个整数数组arr，arr可能 包含重复元素。
每一次操作中，你可以在 arr的任意位置插入任一整数。比方说，如果arr = [1,4,1,2]，那么你可以在中间添加 3得到[1,4,3,1,2]。
你可以在数组最开始或最后面添加整数。

请你返回 最少操作次数，使得target成为arr的一个子序列。

一个数组的 子序列指的是删除原数组的某些元素（可能一个元素都不删除），同时不改变其余元素的相对顺序得到的数组。
	比方说，[2,7,4]是[4,2,3,7,2,1,4]的子序列（加粗元素），但[2,4,2]不是子序列。

思路：
贪心 + （二分查找 实现最长子序列）
*/
func minOperations(target []int, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n) //定义hash表
	for i, val := range target {
		pos[val] = i
	}
	d := []int{} //存放arr的最长公共子序列数组
	for _, val := range arr {
		if idx, has := pos[val]; has {
			//d中二分查找，第一个 大于等于 idx的下标p
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}
