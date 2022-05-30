package 数组

import "sort"

/**
给你一个区间数组 intervals ，其中intervals[i] = [starti, endi] ，且每个starti 都 不同 。
区间 i 的 右侧区间 可以记作区间 j ，并满足 startj>= endi ，且 startj 最小化 。
返回一个由每个区间 i 的 右侧区间 在intervals 中对应下标组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。

思路：切片自定义排序 + 二分查找 T = O(nlogn) T = O(n)
首先我们可以对区间 intervals 的起始位置进行排序，并将每个起始位置 intervals[i][0] 对应的索引 i 存储在数组 Intervals[i][2] 中，
然后枚举每个区间 i 的右端点 intervals[i][1]，利用二分查找来找到大于等于 intervals[i][1] 的最小值 val 即可，
此时区间 i 对应的右侧区间即为左端点为 val 对应的区间的[2]索引。
*/
func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool { //排序后，对于每一个区间就可以用 二分查找了
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	ans := make([]int, n)
	for _, p := range intervals {
		i := sort.Search(n, func(i int) bool {
			return intervals[i][0] >= p[1]
		})
		if i < n {
			ans[p[2]] = intervals[i][2]
		} else {
			ans[p[2]] = -1
		}
	}
	return ans
}
